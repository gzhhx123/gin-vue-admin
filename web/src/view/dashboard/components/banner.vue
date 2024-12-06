<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->

<template>
  <el-carousel class="-mt-2">
    <el-carousel-item
      class="cursor-pointer lg:h-40"
      v-for="(item, index) in banners"
      :key="index"
      @click="openLink(item.link)"
    >
      <el-image class="h-full w-full" :src="item.img" fit="fill"></el-image>
    </el-carousel-item>
  </el-carousel>
</template>

<script setup>
  import banner from '@/assets/banner.jpg'
  import banner2 from '@/assets/banner2.jpg'

  const openLink = (link) => {
    window.open(link, '_blank')
  }

  const banners = [
    {
      img: "https://img2.baidu.com/it/u=257681495,312745373&fm=253&fmt=auto&app=120&f=JPEG?w=750&h=500",
      link: 'https://baidu.com'
    },
    {
      img: "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fimage109.360doc.com%2FDownloadImg%2F2021%2F12%2F2921%2F236804917_4_20211229090912536&refer=http%3A%2F%2Fimage109.360doc.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1736071415&t=a90538677d458e713cc6c3f3fb431d5b",
      link: 'https://baidu.com'
    },
  ]
</script>

<style scoped lang="scss"></style>
