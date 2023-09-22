<template>
  <div>
    <form @submit.prevent="addMarkerAndReturn">
      <input v-model="markerInfo" placeholder="输入标记信息" />
      <button type="submit">确认</button>
      <button @click="cancel">取消</button>
    </form>
  </div>
</template>
  
<script>
export default {
  data() {
    return {
      markerInfo: '' // 用于存储标记的信息
    };
  },
  methods: {
    addMarkerAndReturn() {
  if (this.markerInfo.trim() !== '') {
    // 如果输入了标记信息
    // 创建标记并添加到地图上
    const marker = new this.$parent.$parent.BMap.Marker(this.$parent.$parent.clickedPoint);
    this.$parent.mapInstance.addOverlay(marker);

    // 创建信息窗口显示标记信息
    const infoWindow = new this.$parent.$parent.BMap.InfoWindow(this.markerInfo);
    marker.addEventListener('click', () => {
      this.$parent.$parent.mapInstance.openInfoWindow(infoWindow, this.$parent.$parent.clickedPoint);
    });
  }

  // 返回到map.vue页面
  this.$router.push({ name: 'map' });
},
    cancel() {
      // 取消操作，直接返回到map.vue页面
      this.$router.push({ name: 'map' });
    }
  }
};
</script>
  