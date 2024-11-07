
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="产品:" prop="productId">
        <el-select  v-model="formData.productId" placeholder="请选择产品" style="width:100%" :clearable="true" >
          <el-option v-for="(item,key) in dataSource.productId" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="卖家:" prop="sellerId">
        <el-select  v-model="formData.sellerId" placeholder="请选择卖家" style="width:100%" :clearable="true" >
          <el-option v-for="(item,key) in dataSource.sellerId" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="细节图:" prop="evaluatePics">
           <SelectImage v-model="formData.evaluatePics" multiple file-type="image"/>
       </el-form-item>
        <el-form-item label="估价状态:" prop="status">
           <el-select v-model="formData.status" placeholder="请选择估价状态" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in evaluate_statusOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
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
    getEvaluateDataSource,
  createEvaluate,
  updateEvaluate,
  findEvaluate
} from '@/api/bagique/evaluate'

defineOptions({
    name: 'EvaluateForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
const evaluate_statusOptions = ref([])
const formData = ref({
            productId: undefined,
            sellerId: undefined,
            evaluatePics: [],
            status: '',
            remark: '',
        })
// 验证规则
const rule = reactive({
               productId : [{
                   required: true,
                   message: '请正确选择产品',
                   trigger: ['input','blur'],
               }],
               sellerId : [{
                   required: true,
                   message: '请正确选择卖家',
                   trigger: ['input','blur'],
               }],
               evaluatePics : [{
                   required: true,
                   message: '请正确上传细节图',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '请正确选择估价状态',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getEvaluateDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findEvaluate({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    evaluate_statusOptions.value = await getDictFunc('evaluate_status')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createEvaluate(formData.value)
               break
             case 'update':
               res = await updateEvaluate(formData.value)
               break
             default:
               res = await createEvaluate(formData.value)
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