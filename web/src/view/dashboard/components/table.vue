<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->

<template>
  <div>
    <el-table :data="tableData" stripe style="width: 100%">
      <el-table-column prop="ranking" label="排名" width="80" align="center" />
      <el-table-column prop="title" label="标题" show-overflow-tooltip />
      <el-table-column prop="updated_at" label="更新时间" width="100" align="center" />
    </el-table>
  </div>
</template>

<script setup>
  const tableData = [
    {
      ranking: 1,
      title: '更简洁的使用界面，更快速的操作体验',
      updated_at: '2024-12-06',
      hot: 263
    },
    {
      ranking: 2,
      title: '更优质的服务，更便捷的使用体验',
      updated_at: '2024-12-06',
      hot: 223
    },
    {
      ranking: 3,
      title: '更快速的创意实现，更高效的工作效率',
      updated_at: '2024-12-06',
      hot: 176
    },
    {
      ranking: 4,
      title: '更多的创意资源，更多的创意灵感',
      updated_at: '2024-12-06',
      hot: 145
    },
    {
      ranking: 5,
      title: '更合理的代码结构，更清晰的代码逻辑',
      updated_at: '2024-12-06',
      hot: 110
    }
  ]
</script>

<style scoped lang="scss"></style>
