<template>
	<div id="index" style="width: 100%;">
		<baidu-map class="map" :center="center" :zoom="zoom" :clickedPoint="clickedPoint" @ready="handler" />


		<!-- 窗口 -->
		<el-dialog title="用户登录" :visible.sync="markInfoDialogVisible" class="login-container" :close-on-click-modal="false"
			:center="true">
			<el-input type="textarea" v-model="markInfo" placeholder="请输入内容"></el-input>
			<template v-slot:footer>
				<div>
					<el-row>
						<el-button type="primary" @click="submitMarkInfoForm">确定</el-button>
						<el-button @click="cancelMarkInfoDialog">取消</el-button>
					</el-row>
				</div>
			</template>

		</el-dialog>


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
			clickedPoint: null, // 用于存储点击地图后的坐标
			markInfo: "",
			markInfoDialogVisible: false,
		}
	},
	created() { },
	mounted() { },
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
				console.log(this.clickedPoint.lat, this.clickedPoint.lng)
				this.markInfoDialogVisible = true;
				//this.$router.push({ name: 'post' }); // 导航到Post.vue页面


			})
		},
		cancelMarkInfoDialog() {
			this.markInfoDialogVisible = false;
		},
		submitMarkInfoForm() {
			if (this.markInfo) {
				console.log(this.markInfo);
				this.markInfoDialogVisible = false;
				// 创建标记并添加到地图上
				
				const marker = new BMap.Marker(this.clickedPoint); // 创建标记
				this.mapInstance.addOverlay(marker); // 添加标记到地图
			} else {
				this.$message.error("输入内容不能为空");
			}
		},

	}
}
</script>

<style lang="less" scoped>
.map {
	width: 100%;
	height: 650px;
}
</style>
