<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->

<template>
  <div>
    <el-table :data="tableData" stripe style="width: 100%">
      <el-table-column prop="ranking" label="排名" width="80" align="center" />
      <el-table-column prop="title" label="标题" show-overflow-tooltip>
        <template #default="{ row }">
          <a class="text-active" :href="row.link" target="_blank">{{
            row.title
          }}</a>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script setup>
  const tableData = [
    {
      ranking: 1,
      title: '有新的估价单',
      click_num: 523,
      hot: 263,
      link: 'https://plugin.gin-vue-admin.com/#/layout/newPluginInfo?id=36'
    },
    {
      ranking: 2,
      title:
        '有新的可采购',
      click_num: 416,
      hot: 223,
      link: 'https://plugin.gin-vue-admin.com/#/layout/newPluginInfo?id=42'
    },

  ]
</script>

<style scoped lang="scss"></style>
