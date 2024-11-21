package com.example.android.adapter


import android.view.View
import android.widget.TextView
import androidx.recyclerview.widget.RecyclerView
import com.example.android.R

class ViewHolderList (item: View) : RecyclerView.ViewHolder(item) {

    val characterList: TextView = item.findViewById(R.id.Text1)
}
