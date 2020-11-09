package com.inkeliz.giohyperlink_android;

import android.app.Activity;
import android.view.View;
import android.content.Context;
import android.content.Intent;
import android.net.Uri;

public class giohyperlink_android {

    public void open(View view, String url) {

        // Create the Intent, which is to open the given URL
        Intent intent = new Intent(Intent.ACTION_VIEW, Uri.parse(url));

        // Get GioActivity from GioView
        Activity activity = (Activity)view.getContext();

        // Run on main thread, from GioActivity
        activity.runOnUiThread(new Runnable() {
            public void run() {

                // Create the activity from GioActivity
                activity.startActivity(intent);
            }
        });

    }

}