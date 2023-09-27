<template>
	<div id="index" style="width: 100%;">
		<baidu-map class="map" :center="center" :zoom="zoom" :clickedPoint="clickedPoint" @ready="handler">
			<bm-geolocation anchor="BMAP_ANCHOR_BOTTOM_RIGHT" :showAddressBar="true" :autoLocation="true"
				@locationSuccess="handleLocationSuccess"></bm-geolocation>
		</baidu-map>

		<!-- 窗口 -->
		<el-dialog title="帖子详情" :visible.sync="markInfoDialogVisible" class="login-container" :close-on-click-modal="false"
			:center="true">
			<el-form :model="detailForm" label-width="80px" ref="detailForm">
				<el-form-item label="地址:">
					{{ this.address }}
				</el-form-item>
				<el-form-item label="标题:">
					<el-input v-model="detailForm.title" placeholder="请输入标题"></el-input>
				</el-form-item>
				<el-form-item label="内容:">
					<el-input type="textarea" :rows="5" v-model="detailForm.detail" placeholder="请输入内容"
						auto-complete="off"></el-input>
				</el-form-item>
				<el-form-item label="图片:">
					<el-upload class="upload-demo" action="https://jsonplaceholder.typicode.com/posts/"
						:on-preview="handlePreview" :on-remove="handleRemove" :file-list="fileList" list-type="picture">
						<el-button size="small" type="primary">点击上传</el-button>
						<div class="el-upload__tip">只能上传jpg/png文件,且不超过500kb</div>
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
				pictures: "",
			},
			"Location": {
				city: "",
				district: "",
				province: "",
				street: "",
				streetNumber: "",
			},
			"Pictures": "",
			fileList: [{ name: 'food.jpeg', url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100' }, { name: 'food2.jpeg', url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100' }],
			markers: [],
		}
	},
	created() { },
	mounted() { },
	methods: {
		handler({ BMap, map }) {
			//console.log(BMap, map)
			this.center.lng = 114.365777	//115.3542
			this.center.lat = 30.531512		//30.5286

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
				// 创建地理编码器实例
				const geocoder = new BMap.Geocoder();

				// 执行逆地理编码，获取地址信息
				geocoder.getLocation(new BMap.Point(lng, lat), (result) => {
					if (result) {
						// 在这里处理地址信息，例如将地址信息显示在页面上
						this.address = result.address; // 完整地址
						this.province = result.addressComponents.province; // 省份
						this.city = result.addressComponents.city; // 城市
						this.district = result.addressComponents.district; // 区域
						this.street = result.addressComponents.street; // 街道
						this.streetNumber = result.addressComponents.streetNumber; // 街道编号

						// console.log("地址:", this.address);
						// console.log("省份:", this.province);
						// console.log("城市:", this.city);
						// console.log("区域:", this.district);
						// console.log("街道:", this.street);
						// console.log("街道编号:", this.streetNumber);

						// 这里可以将地址信息展示在页面上或进行其他处理
					} else {
						console.log("无法获取地址信息");
					}
				});
				this.markInfoDialogVisible = true;
			})

			//绘制用户已经保存过的点
			let length = this.$store.state.markedPoints.length;
			//console.log(this.$store.state.markedPoints);
			for (var i = 0; i < length; ++i) {
				this.markers.push(new BMap.Marker(new BMap.Point(this.$store.state.markedPoints[i].point.lng, this.$store.state.markedPoints[i].point.lat)));
				//console.log(BMap.Marker);

				this.mapInstance.addOverlay(this.markers[i]);
				const vm = this;
				//console.log(this.markers[i])

				const temp = i;
				this.markers[i].addEventListener('click', function () {
					console.log(vm.$store.state.markedPoints[temp].id);
					vm.$router.push('/detail');
				});
			}

		},
		cancelMarkInfoDialog() {
			this.markInfoDialogVisible = false;
		},
		submitMarkInfoForm() {
			if (this.detailForm.title && this.detailForm.detail) {
				console.log(this.detailForm.title, this.detailForm.detail);
				this.markInfoDialogVisible = false;

				// 创建标记并添加到地图上
				console.log(this.clickedPoint)
				const marker = new BMap.Marker(this.clickedPoint); // 创建标记
				this.mapInstance.addOverlay(marker);
				//console.log(marker); // 添加标记到地图
				const vm = this
				marker.addEventListener('click', function () {
					vm.$router.push("/detail");
					//alert('你单击了标记');
				});

				//后端交互请求在这里写
				// axios.post("/api/article", {
				// 	"Title": "hello",
				// 	"Detail": "klwajdlawdkawdhwadawhjkbgdhjvvbfv",
				// 	"lat": 111,
				// 	"lng": 111,
				// 	"Location": {
				// 		"city": "1",
				// 		"district": "1",
				// 		"province": "11",
				// 		"street": "111",
				// 		"streetNumber": "111"
				// 	},
				// 	"Cover": "lhkelk"
				// }, {
				// 	headers: {
				// 		Authorization: "Bearer" + this.$store.state.token	 // 在请求头中添加 Authorization 字段，值为 Bearer + 空格 + token
				// 	}
				// })
				// 	.then(response => {
				// 		// 处理响应
				// 	})
				// 	.catch(error => {
				// 		// 处理错误
				// 	});

				this.detailForm.title = "";
				this.detailForm.detail = "";
			} else {
				this.$message.error("标题与内容不能为空");
			}
		},
		handleRemove(file, fileList) {
			//console.log(file, fileList);
		},
		handlePreview(file) {
			//console.log(file);
		},
		handleLocationSuccess(event) {
			this.address = event.addressComponent; // 获取地址信息
			console.log(this.address);
			//console.log(event);

		},
		// submitLoginForm() {   //登录逻辑
		// 	if (this.loginForm.username && this.loginForm.password) {
		// 		axios.post("/", {
		// 			"UserName": this.$store.state.username,
		// 			"Title": this.detailForm.title,
		// 			"Detail": this.detailForm.detail,
		// 			"Lat": this.clickedPoint.lat,
		// 			"Lng": this.clickedPoint.lng,
		// 			"Location": {
		// 				city: this.city,
		// 				district: this.district,
		// 				province: this.province,
		// 				street: this.street,
		// 				streetNumber: this.streetNumber,
		// 			},
		// 			"Pictures": ""
		// 		})
		// 			.then(response => {
		// 				// 登录成功处理逻辑
		// 				if (response.data.code === 20000) {
		// 					// 更新登录状态
		// 					this.$store.commit("isLogInTrue");
		// 					// 跳转到主页或其他页面
		// 					this.closeLoginDialog();
		// 					this.$router.push("/");
		// 					this.$message.success("登录成功！");

		// 					this.$store.commit("changeUserName", this.loginForm.username);
		// 				}
		// 			})
		// 			.catch(error => {
		// 				// 登录失败处理逻辑
		// 				this.loginForm.errorMessage = "登录失败，请重试";
		// 				this.$message.error(this.loginForm.errorMessage);
		// 				this.loginForm.password = "";
		// 			});
		// 	} else {
		// 		this.loginForm.errorMessage = "用户名或密码不能为空"
		// 		this.$message.error(this.loginForm.errorMessage);
		// 		this.loginForm.password = "";
		// 	}
		// },

	},

}

</script>

<style lang="less" scoped>
.map {
	width: 100%;
	height: 650px;
}
</style>
