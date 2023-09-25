<template>
	<div id="index" style="width: 100%;">
		<baidu-map class="map" :center="center" :zoom="zoom" :clickedPoint="clickedPoint" @ready="handler">
			<bm-geolocation anchor="BMAP_ANCHOR_BOTTOM_RIGHT" :showAddressBar="true" :autoLocation="true" @locationSuccess="handleLocationSuccess"></bm-geolocation>
		</baidu-map>

		<!-- 窗口 -->
		<el-dialog title="帖子详情" :visible.sync="markInfoDialogVisible" class="login-container" :close-on-click-modal="false"
			:center="true">
			<el-form :model="detailForm" label-width="80px" ref="detailForm">
				<el-form-item label="标题">
					<el-input v-model="detailForm.title" placeholder="请输入标题"></el-input>
				</el-form-item>
				<el-form-item label="内容">
					<el-input type="textarea" :rows="5" v-model="detailForm.detail" placeholder="请输入内容"
						auto-complete="off"></el-input>
				</el-form-item>
				<el-form-item label="图片">
					<el-upload class="upload-demo" action="https://jsonplaceholder.typicode.com/posts/"
						:on-preview="handlePreview" :on-remove="handleRemove" :file-list="fileList" list-type="picture">
						<el-button size="small" type="primary">点击上传</el-button>
						<div slot="tip" class="el-upload__tip">只能上传jpg/png文件,且不超过500kb</div>
					</el-upload>
				</el-form-item>
			</el-form>
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
			address: "",
			markInfoDialogVisible: false,
			detailForm: {
				title: "",
				detail: "",
				pictures: [],
			},
			fileList: [{ name: 'food.jpeg', url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100' }, { name: 'food2.jpeg', url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100' }]
		}
	},
	created() { },
	mounted() { },
	methods: {
		handler({ BMap, map }) {
			console.log(BMap, map)
			this.center.lng = 115.3542
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
			this.mapInstance.addControl(new BMap.CityListControl({ anchor: BMAP_ANCHOR_TOP_RIGHT }));

			// 添加点击事件监听器
			this.mapInstance.addEventListener('click', (e) => {

				const { lng, lat } = e.point; // 获取点击位置的经纬度
				this.clickedPoint = e.point; // 将坐标保存到clickedPoint
				console.log(this.clickedPoint.lat, this.clickedPoint.lng)
				this.markInfoDialogVisible = true;


			})
		},
		cancelMarkInfoDialog() {
			this.markInfoDialogVisible = false;
		},
		submitMarkInfoForm() {
			if (this.detailForm.title && this.detailForm.detail) {
				console.log(this.detailForm.title, this.detailForm.detail);
				this.markInfoDialogVisible = false;
				// 创建标记并添加到地图上

				const marker = new BMap.Marker(this.clickedPoint); // 创建标记
				this.mapInstance.addOverlay(marker); // 添加标记到地图
			} else {
				this.$message.error("标题与内容不能为空");
			}
		},
		handleRemove(file, fileList) {
			console.log(file, fileList);
		},
		handlePreview(file) {
			console.log(file);
		},
		handleLocationSuccess(event) {
			this.address = event.addressComponent; // 获取地址信息
			console.log(this.address);
			//console.log(event);

		},

	},

}

</script>

<style lang="less" scoped>
.map {
	width: 100%;
	height: 650px;
}
</style>
