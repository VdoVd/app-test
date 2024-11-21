package com.example.android

import android.os.Bundle
import android.webkit.WebSettings
import android.webkit.WebView
import android.webkit.WebViewClient
import androidx.appcompat.app.AppCompatActivity


//import com.example.android.Fragment.Home


class MainActivity : AppCompatActivity() {

    private lateinit var webView: WebView
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
        webView = findViewById(R.id.webView)
        webView.webViewClient= WebViewClient()
        webView.loadUrl("file:///android_asset/index.html")
        webView.settings.javaScriptEnabled=true
        webView.settings.setSupportZoom(true)
    }

    // if you press Back button this code will work



}

