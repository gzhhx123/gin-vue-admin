
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="估价公司:" prop="companyId">
        <el-select  v-model="formData.companyId" placeholder="请选择估价公司" style="width:100%" :clearable="true" >
          <el-option v-for="(item,key) in dataSource.companyId" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="估价记录id:" prop="evaluateId">
          <el-input v-model.number="formData.evaluateId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="估价:" prop="price">
          <el-input-number v-model="formData.price" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="估价成本:" prop="fee">
          <el-input-number v-model="formData.fee" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="备注:" prop="remark">
          <el-input v-model="formData.remark" :clearable="true"  placeholder="请输入备注" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
    getEvaluatePriceDataSource,
  createEvaluatePrice,
  updateEvaluatePrice,
  findEvaluatePrice
} from '@/api/bagique/evaluatePrice'

defineOptions({
    name: 'EvaluatePriceForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            companyId: undefined,
            evaluateId: undefined,
            price: 0,
            fee: 0,
            remark: '',
        })
// 验证规则
const rule = reactive({
               companyId : [{
                   required: true,
                   message: '请正确选择估价公司',
                   trigger: ['input','blur'],
               }],
               evaluateId : [{
                   required: true,
                   message: '请正确选择估价记录id',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getEvaluatePriceDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findEvaluatePrice({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createEvaluatePrice(formData.value)
               break
             case 'update':
               res = await updateEvaluatePrice(formData.value)
               break
             default:
               res = await createEvaluatePrice(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>