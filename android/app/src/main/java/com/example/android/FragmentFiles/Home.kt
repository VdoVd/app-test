package com.example.android.FragmentFiles

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.fragment.app.Fragment
import androidx.recyclerview.widget.LinearLayoutManager
import androidx.recyclerview.widget.RecyclerView


import com.example.android.R
import com.example.android.adapter.RecyclerAdapter

class Home :Fragment() {

    private lateinit var recyclerView: RecyclerView
    private lateinit var adapter:RecyclerAdapter
    private val data = listOf(
        "マリオ",
        "ルイージ",
        "ピーチ",
        "クッパ",
        "デイジー",
        "ワリオ",
        "ワルイージ",
        "ヘイホー",
        "マリオ",
        "ルイージ",
        "ピーチ",
        "クッパ",
        "デイジー",
        "ワリオ",
        "ワルイージ",
        "ヘイホー",  "マリオ",
        "ルイージ",
        "ピーチ",
        "クッパ",
        "デイジー",
        "ワリオ",
        "ワルイージ",
        "ヘイホー",  "マリオ",
        "ルイージ",
        "ピーチ",
        "クッパ",
        "デイジー",
        "ワリオ",
        "ワルイージ",
        "ヘイホー"
    )

    override fun onCreateView(inflater: LayoutInflater, container: ViewGroup?, savedInstanceState: Bundle?):
            View? {
        val view = inflater.inflate(R.layout.home, container, false)
        recyclerView = view.findViewById(R.id.recyclerList)
        adapter = RecyclerAdapter(data)
        recyclerView.adapter = adapter
        recyclerView.layoutManager = LinearLayoutManager(this.context)
        return recyclerView
    }


}
