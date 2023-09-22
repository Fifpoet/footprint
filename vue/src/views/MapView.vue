<template>
	<div id="index" style="width: 100%;">
		<baidu-map class="map" :center="center" :zoom="zoom" :clickedPoint="clickedPoint" @ready="handler" />
	</div>
</template>

<script>
export default {
	name: 'Index',
	components: {},
	data() {
		return {
			center: { lng: 0, lat: 0 },
			zoom: 0,
      clickedPoint: null // 用于存储点击地图后的坐标
		}
	},
	created() {},
	mounted() {},
	methods: {
		handler({ BMap, map }) {
			console.log(BMap, map)
			this.center.lng = 114.3542
			this.center.lat = 30.5286
			this.zoom = 15
      // 获取地图实例
			this.mapInstance = map;
			
			// 启用鼠标滚轮缩放
			this.mapInstance.enableScrollWheelZoom(true);
      // 添加地图控件
      this.mapInstance.addControl(new BMap.NavigationControl());
      this.mapInstance.addControl(new BMap.ScaleControl());
      this.mapInstance.addControl(new BMap.OverviewMapControl());

      // 添加点击事件监听器
      this.mapInstance.addEventListener('click', (e) => {
        const { lng, lat } = e.point; // 获取点击位置的经纬度
        this.clickedPoint = e.point; // 将坐标保存到clickedPoint
        this.$router.push({ name: 'post' }); // 导航到Post.vue页面
       
		  })
  }
    
	}
}
</script>

<style lang="less" scoped>
.map{
	width: 100%;
	height: 650px;
}
</style>
