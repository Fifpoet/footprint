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
					<el-upload class="upload-demo" action="https://jsonplaceholder.typicode.com/posts/" :on-preview="handlePreview" :on-remove="handleRemove" :file-list="fileList" list-type="picture"
						:file-size-limit="fileSizeLimit">
						<el-button size="small" type="primary">点击上传</el-button>
						<div class="el-upload__tip">只能上传jpg/png文件,且不超过3MB</div>
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
				picture: "",
			},
			"Location": {
				city: "",
				district: "",
				province: "",
				street: "",
				streetNumber: "",
			},
			// fileList: [{
			// 	name: "Test1.jpg",
			// 	url: "http://127.0.0.1:9000/footprint01/%E6%B5%8B%E8%AF%951.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=1KN7TBQEHZADZEZIOW32%2F20230928%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20230928T073438Z&X-Amz-Expires=604800&X-Amz-Security-Token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NLZXkiOiIxS043VEJRRUhaQURaRVpJT1czMiIsImV4cCI6MTY5NTkyNjg1NSwicGFyZW50IjoibWluaW9hZG1pbiJ9.V2ZoRuwWEUUPS7VVb7r_goi3NgpCU_EiNsdzPiiOgv8nqNWzvMuILiM1JV71abvasGm56c7_lVixyRlRrgsfwQ&X-Amz-SignedHeaders=host&versionId=null&X-Amz-Signature=ddc4f35aeceaa4a7fc09016947051c13378212808642fbca19a555795f56e873"
			// }],
			markers: [],
			fileSizeLimit: 3 * 1024 * 1024, // 文件大小限制为 1MB
		}
	},
	created() { },
	mounted() { },
	methods: {
		handler({ BMap, map }) {
			//console.log(BMap, map)
			this.center.lng = 114.365777
			this.center.lat = 30.531512

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
				this.mapInstance.addOverlay(this.markers[i]);
				const vm = this;

				const temp = i;
				this.markers[i].addEventListener('click', function () {
					console.log(vm.$store.state.markedPoints[temp].id);
					vm.$store.commit("changePostDetail", vm.$store.state.markedPoints[temp]);
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

				// axios.post("/api/upload", this.Picture, {
				// 	headers: {
				// 		Authorization: "Bearer " + this.$store.state.token
				// 		// 在请求头中添加 Authorization 字段，值为 Bearer + 空格 + token
				// 	}

				// })

				//后端交互请求在这里写
				axios.post("/api/article", {
					"Title": this.detailForm.title,
					"Detail": this.detailForm.detail,
					"lng": this.clickedPoint.lng,
					"lat": this.clickedPoint.lat,
					"Location": {
						"city": this.city,
						"district": this.district,
						"province": this.province,
						"street": this.street,
						"streetNumber": this.streetNumber,
					},
					"Cover": "http://127.0.0.1:9000/footprint01/Test1.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=1KN7TBQEHZADZEZIOW32%2F20230928%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20230928T075140Z&X-Amz-Expires=604800&X-Amz-Security-Token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NLZXkiOiIxS043VEJRRUhaQURaRVpJT1czMiIsImV4cCI6MTY5NTkyNjg1NSwicGFyZW50IjoibWluaW9hZG1pbiJ9.V2ZoRuwWEUUPS7VVb7r_goi3NgpCU_EiNsdzPiiOgv8nqNWzvMuILiM1JV71abvasGm56c7_lVixyRlRrgsfwQ&X-Amz-SignedHeaders=host&versionId=null&X-Amz-Signature=9116ddd3fab6378e239100064f4242abd7410a0dccddc7c61629e75c0bb0aa29"
				}, {
					headers: {
						Authorization: "Bearer " + this.$store.state.token
						// 在请求头中添加 Authorization 字段，值为 Bearer + 空格 + token
					}
				})
					.then(response => {
						// 处理响应
						this.$message.success("上传成功")
					})
					.catch(error => {
						// 处理错误
						this.$message.error("上传失败")
					});

				this.detailForm.title = "";
				this.detailForm.detail = "";

				this.$store.commit("clearMarkedPoint");

				axios.get("/api/article", {
					headers: {
						Authorization: "Bearer " + this.$store.state.token
						// 在请求头中添加 Authorization 字段，值为 Bearer + 空格 + token
					}
				})
					.then(response => {
						// 处理成功响应
						console.log(response.data.articles); // 输出获取到的 JSON 数据

						for (var item of response.data.articles) {
							let markedPoint = {
								point: {
									lng: item.Lng,
									lat: item.Lat,
								},
								id: item.ID,
								title: item.Title,
								detail: item.Content,
								pictures: "",
								location: {
									city: item.City,
									district: item.District,
									province: item.Province,
								},
								cover: item.cover,
							}
							this.$store.commit("addMarkedPoint", markedPoint);

						}
					})
					.catch(error => {
						// 处理错误响应
						console.error(error);
					});

			} else {
				this.$message.error("标题与内容不能为空");
			}
		},
		handleRemove(file, fileList) {
			console.log(file, fileList);
		},
		handlePreview(file) {
			// 预览文件的回调函数
			console.log(file); // 要预览的文件对象
			// 在这里可以编写文件的预览逻辑，例如打开新的窗口显示文件内容等
			window.open(file.url); // 示例：打开新窗口显示文件内容
		},
		handleLocationSuccess(event) {
			this.address = event.addressComponent; // 获取地址信息
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
