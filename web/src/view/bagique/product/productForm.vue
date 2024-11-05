
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="品牌:" prop="brandId">
        <el-select  v-model="formData.brandId" placeholder="请选择品牌" style="width:100%" :clearable="true" >
          <el-option v-for="(item,key) in dataSource.brandId" :key="key" :label="item.label" :value="item.value" />
        </el-select>
       </el-form-item>
        <el-form-item label="产品名称:" prop="productName">
          <el-input v-model="formData.productName" :clearable="true"  placeholder="请输入产品名称" />
       </el-form-item>
        <el-form-item label="产品sku:" prop="productSku">
          <el-input v-model="formData.productSku" :clearable="true"  placeholder="请输入产品sku" />
       </el-form-item>
        <el-form-item label="参考价（Min）:" prop="referPriceMin">
          <el-input-number v-model="formData.referPriceMin" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="参考价（Max）:" prop="referPriceMax">
          <el-input-number v-model="formData.referPriceMax" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="参考图片:" prop="referPics">
           <SelectImage v-model="formData.referPics" multiple file-type="image"/>
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
    getProductDataSource,
  createProduct,
  updateProduct,
  findProduct
} from '@/api/bagique/product'

defineOptions({
    name: 'ProductForm'
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
const formData = ref({
            brandId: undefined,
            productName: '',
            productSku: '',
            referPriceMin: 0,
            referPriceMax: 0,
            referPics: [],
            remark: '',
        })
// 验证规则
const rule = reactive({
               productName : [{
                   required: true,
                   message: '请正确填写产品名称',
                   trigger: ['input','blur'],
               }],
               productSku : [{
                   required: true,
                   message: '请正确填写产品sku',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getProductDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findProduct({ ID: route.query.id })
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
               res = await createProduct(formData.value)
               break
             case 'update':
               res = await updateProduct(formData.value)
               break
             default:
               res = await createProduct(formData.value)
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