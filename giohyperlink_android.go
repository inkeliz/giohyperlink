// +build android

package giohyperlink

import (
	"gioui.org/app"
	"gioui.org/io/event"
	"git.wow.st/gmp/jni"
	"net/url"
)

//go:generate javac -source 8 -target 8 -bootclasspath $ANDROID_HOME\platforms\android-29\android.jar -d $TEMP\giohyperlink\classes giohyperlink_android.java
//go:generate jar cf giohyperlink_android.jar -C $TEMP\giohyperlink\classes .

var view uintptr

func listenEvents(event event.Event) {
	if e, ok := event.(app.ViewEvent); ok {
		view = e.View
	}
}

func open(u *url.URL) error {
	if view == 0 {
		return ErrNotReady
	}

	return jni.Do(jni.JVMFor(app.JavaVM()), func(env jni.Env) error {

		// Get the GioView object
		obj := jni.Object(view)
		cls := jni.GetObjectClass(env, obj)

		// Run getClass() to get the Class of GioView
		mid := jni.GetMethodID(env, cls, "getClass", "()Ljava/lang/Class;")
		obj, err := jni.CallObjectMethod(env, obj, mid)
		if err != nil {
			panic(err)
		}

		// Run getClassLoader() to get the ClassLoader from Class
		cls = jni.GetObjectClass(env, obj)
		mid = jni.GetMethodID(env, cls, "getClassLoader", "()Ljava/lang/ClassLoader;")
		obj, err = jni.CallObjectMethod(env, obj, mid)
		if err != nil {
			panic(err)
		}

		// Run findClass() from ClassLoader. The return is our custom class (in that case it's the
		// com.inkeliz.giohyperlink.giohyperlink_android, that name is defined on `giohyperlink_android.java`.
		cls = jni.GetObjectClass(env, obj)
		mid = jni.GetMethodID(env, cls, "findClass", "(Ljava/lang/String;)Ljava/lang/Class;")
		clso, err := jni.CallObjectMethod(env, obj, mid, jni.Value(jni.JavaString(env, `com.inkeliz.giohyperlink.giohyperlink_android`)))
		if err != nil {
			panic(err)
		}

		// We need to convert Object (because we use CallObjectMethod) to jni.Class type
		cls = jni.Class(clso)

		// Create a new Object from our class. It's almost the same of `new giohyperlink_android()`
		// The `<init>` and `NewObject` are used to create a "variable" with the class that we get before.
		mid = jni.GetMethodID(env, cls, "<init>", `()V`)
		obj, err = jni.NewObject(env, cls, mid)
		if err != nil {
			panic(err)
		}

		// Run the `open()` function from our custom Java. That is defined inside the `giohyperlink_android` class
		// you can view that at giohyperlink_android.java.
		//
		// Our java function is:
		// public void open(View view, String url) {}
		//
		// So we need to supply the view argument and the url argument. That view argument is the `GioView` itself
		// the GioView is the `view` variable, which we got from app.ViewEvents.
		mid = jni.GetMethodID(env, cls, "open", "(Landroid/view/View;Ljava/lang/String;)V")
		err = jni.CallVoidMethod(env, obj, mid, jni.Value(view), jni.Value(jni.JavaString(env, u.String())))
		if err != nil {
			panic(err)
		}

		return nil
	})
}
