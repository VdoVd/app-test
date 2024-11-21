<template>
	<view class="content">
		<view class="top">
			<view>
				<img :src="avatar"/>
				<view>{{user==undefined?"登录":"user.name"}}</view>
			</view>
			<view>
				<view>
					
				</view>
				<view>
					
				</view>
				<view>
					
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	import avatar from '@/static/mine.svg'
	export default {
		data() {
			return {
				title: 'Hello',
				user:undefined
			}
		},
		onLoad() {

		},
		methods: {

		}
	}
</script>

<style>
	.content {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
	}

	.logo {
		height: 200rpx;
		width: 200rpx;
		margin-top: 200rpx;
		margin-left: auto;
		margin-right: auto;
		margin-bottom: 50rpx;
	}

	.text-area {
		display: flex;
		justify-content: center;
	}

	.title {
		font-size: 36rpx;
		color: #8f8f94;
	}
</style>
