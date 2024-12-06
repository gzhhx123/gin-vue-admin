<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->

<template>
  <el-scrollbar>
    <div
      v-for="(item, index) in notices"
      :key="index"
      class="flex items-center mb-1.5 gap-3"
    >
      <el-tag :type="item.type" size="small">
        {{ item.typeTitle }}
      </el-tag>
      <el-tooltip effect="light" :content="item.title" placement="top">
        <div class="text-xs text-gray-700 dark:text-gray-300 line-clamp-1">
          {{ item.title }}
        </div>
      </el-tooltip>
    </div>
  </el-scrollbar>
</template>

<script setup>
  const notices = [
    {
      type: 'primary',
      typeTitle: '公告',
      title: '测试公告1'
    },
    {
      type: 'success',
      typeTitle: '通知',
      title: '测试通知1'
    },
    {
      type: 'warning',
      typeTitle: '警告',
      title: '测试警告1'
    },
    {
      type: 'danger',
      typeTitle: '违规',
      title: '测试违规1'
    },
    {
      type: 'info',
      typeTitle: '信息',
      title: '测试信息1'
    },
    {
      type: 'primary',
      typeTitle: '公告',
      title: '测试公告2'
    },
    {
      type: 'success',
      typeTitle: '通知',
      title: '测试通知2'
    },
    {
      type: 'warning',
      typeTitle: '警告',
      title: '测试警告2'
    },
    {
      type: 'danger',
      typeTitle: '错误',
      title: '测试错误2'
    },
    {
      type: 'info',
      typeTitle: '信息',
      title: '测试信息2'
    }
  ]
</script>

<style scoped lang="scss"></style>
