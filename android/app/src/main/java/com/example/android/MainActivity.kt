package com.example.android

import android.os.Bundle

import com.example.android.adapter.MyPagerAdapter
import com.google.android.material.tabs.TabLayout
import androidx.appcompat.app.AppCompatActivity
import androidx.appcompat.widget.Toolbar
import androidx.recyclerview.widget.LinearLayoutManager
import androidx.recyclerview.widget.RecyclerView
import androidx.viewpager.widget.ViewPager
import com.example.android.FragmentFiles.ActivityFragment
import com.example.android.FragmentFiles.Find
import com.example.android.FragmentFiles.Home
import com.example.android.FragmentFiles.Mine
import com.example.android.adapter.RecyclerAdapter

//import com.example.android.Fragment.Home


class MainActivity : AppCompatActivity() {
    private lateinit var pager: ViewPager // creating object of ViewPager
    private lateinit var tab: TabLayout  // creating object of TabLayout
    private lateinit var bar: Toolbar    // creating object of ToolBar
    private lateinit var recyclerView: RecyclerView

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        pager = findViewById(R.id.viewPager)
        tab = findViewById(R.id.tabs)
//        bar = findViewById(R.id.toolbar)
//        setSupportActionBar(bar)

        // Initializing the ViewPagerAdapter
        val adapter = MyPagerAdapter(supportFragmentManager)
        adapter.addFragment(Home(),"首页")
        adapter.addFragment(Find(),"发现")
        adapter.addFragment(ActivityFragment(),"活动")
        adapter.addFragment(Mine(),"我的")
        pager.adapter = adapter
        tab.setupWithViewPager(pager)
    }




}
