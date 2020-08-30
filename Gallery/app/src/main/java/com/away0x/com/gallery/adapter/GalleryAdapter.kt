package com.away0x.com.gallery.adapter

import android.annotation.SuppressLint
import android.graphics.drawable.Drawable
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.navigation.findNavController
import androidx.recyclerview.widget.DiffUtil
import androidx.recyclerview.widget.ListAdapter
import androidx.recyclerview.widget.RecyclerView
import androidx.recyclerview.widget.StaggeredGridLayoutManager
import androidx.viewbinding.ViewBinding
import com.away0x.com.gallery.R
import com.away0x.com.gallery.databinding.GalleryCellBinding
import com.away0x.com.gallery.databinding.GalleryFooterBinding
import com.away0x.com.gallery.fragment.GalleryFragmentDirections
import com.away0x.com.gallery.model.PhotoItem
import com.away0x.com.gallery.viewModel.DATA_STATUS_CAN_LOAD_MORE
import com.away0x.com.gallery.viewModel.DATA_STATUS_NETWORK_ERROR
import com.away0x.com.gallery.viewModel.DATA_STATUS_NO_MORE
import com.away0x.com.gallery.viewModel.GalleryViewModel
import com.bumptech.glide.Glide
import com.bumptech.glide.load.DataSource
import com.bumptech.glide.load.engine.GlideException
import com.bumptech.glide.request.RequestListener
import com.bumptech.glide.request.target.Target

class GalleryAdapter(private val galleryViewModel: GalleryViewModel) : ListAdapter<PhotoItem, MyViewHolder>(DiffCallback) {

    companion object {
        const val NORMAL_VIEW_TYPE = 0
        const val FOOTER_VIEW_TYPE = 1
    }
    // footer 状态
    var footerViewStatus = DATA_STATUS_CAN_LOAD_MORE

    override fun getItemCount() = super.getItemCount() + 1 // + 1 是为了显示 footer

    // 最后一行返回 footer
    override fun getItemViewType(position: Int) = if (isFooter(position)) FOOTER_VIEW_TYPE else NORMAL_VIEW_TYPE

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): MyViewHolder {
        // 是 footer
        if (viewType == FOOTER_VIEW_TYPE) {
            val binding = GalleryFooterBinding.inflate(LayoutInflater.from(parent.context), parent, false).also {
                (it.root.layoutParams as StaggeredGridLayoutManager.LayoutParams).isFullSpan = true // 占整个空间
                it.root.setOnClickListener {_ ->
                    it.progressBar.visibility = View.VISIBLE
                    it.textView.text = "正在加载"
                    galleryViewModel.fetchData()
                }
            }
            return MyViewHolder(binding)
        }

        val binding = GalleryCellBinding.inflate(LayoutInflater.from(parent.context), parent, false)
        val holder = MyViewHolder(binding)
        // 跳转 photo 页
        holder.itemView.setOnClickListener {
            val action = GalleryFragmentDirections.actionGalleryFragmentToPhotoFragment(
                currentList.toTypedArray(),
                holder.adapterPosition
            )
            holder.itemView.findNavController().navigate(action)
        }

        return holder
    }

    @SuppressLint("CheckResult")
    override fun onBindViewHolder(holder: MyViewHolder, position: Int) {
        // 是 footer
        if (isFooter(position)) {
            with(holder.binding as GalleryFooterBinding) {
                when (footerViewStatus) {
                    DATA_STATUS_CAN_LOAD_MORE -> {
                        progressBar.visibility = View.VISIBLE
                        textView.text = "正在加载"
                        root.isClickable = false
                    }
                    DATA_STATUS_NO_MORE -> {
                        progressBar.visibility = View.GONE
                        textView.text = "全部加载完毕"
                        root.isClickable = false
                    }
                    DATA_STATUS_NETWORK_ERROR -> {
                        progressBar.visibility = View.GONE
                        textView.text = "网络故障，点击重试"
                        root.isClickable = true
                    }
                }
            }
            return
        }

        val photoItem = getItem(position)

        with(holder.binding as GalleryCellBinding) {
            shimmerLayoutCell.apply {
                setShimmerColor(0x55FFFFFF)
                setShimmerAngle(0)
                startShimmerAnimation()
            }

            textViewUser.text = photoItem.photoUser
            textViewLikes.text = photoItem.photoLikes.toString()
            textViewFavorites.text = photoItem.photoFavorites.toString()

            // item 确定高度后，瀑布流布局就不会因为计算高度而产生重排
            imageView.layoutParams.height = photoItem.photoHeight
        }

        Glide.with(holder.binding.root)
            .load(photoItem.previewUrl)
            // .placeholder(R.drawable.ic_photo_gray_24dp)
            .placeholder(R.drawable.photo_placeholder)
            .listener(object : RequestListener<Drawable> {
                override fun onLoadFailed(
                    e: GlideException?,
                    model: Any?,
                    target: Target<Drawable>?,
                    isFirstResource: Boolean
                ): Boolean = false

                override fun onResourceReady(
                    resource: Drawable?,
                    model: Any?,
                    target: Target<Drawable>?,
                    dataSource: DataSource?,
                    isFirstResource: Boolean
                ): Boolean {
                    return false.also {
                        holder.binding.shimmerLayoutCell?.stopShimmerAnimation()
                    }
                }
            })
            .into(holder.binding.imageView)
    }

    // 列表数据的差异化比对是在后台异步执行的
    object DiffCallback : DiffUtil.ItemCallback<PhotoItem>() {
        override fun areItemsTheSame(oldItem: PhotoItem, newItem: PhotoItem) = oldItem.photoId == newItem.photoId
        override fun areContentsTheSame(oldItem: PhotoItem, newItem: PhotoItem) = oldItem == newItem
    }

    private fun isFooter(position: Int) = position == itemCount - 1
}

class MyViewHolder(val binding: ViewBinding) : RecyclerView.ViewHolder(binding.root)
