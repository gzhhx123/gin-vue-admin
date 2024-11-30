
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
      
            <el-form-item label="品牌" prop="brandId">
               <el-select  v-model="searchInfo.brandId" placeholder="请选择品牌" :clearable="true" >
                   <el-option v-for="(item,key) in dataSource.brandId" :key="key" :label="item.label" :value="item.value" />
               </el-select>
            </el-form-item>
        <el-form-item label="产品名称" prop="productName">
         <el-input v-model="searchInfo.productName" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="产品sku" prop="productSku">
         <el-input v-model="searchInfo.productSku" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="备注" prop="remark">
         <el-input v-model="searchInfo.remark" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="显示已删除" prop="isRemove">
          <el-switch v-model="searchInfo.isRemove" />
        </el-form-item>
        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
        <div>实时汇率:{{rate.rate}} 更新时间:{{formatDate(rate.time)}}</div>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            <ExportTemplate  template-id="bagique_Product" />
            <ExportExcel  template-id="bagique_Product" />
            <ImportExcel  template-id="bagique_Product" @on-success="getTableData" />
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column sortable align="left" label="创建日期" prop="CreatedAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
        <el-table-column sortable align="left" label="品牌" prop="brandId" width="120">
          <template #default="scope">
                
                    <span>{{ filterDataSource(dataSource.brandId,scope.row.brandId) }}</span>
                
         </template>
         </el-table-column>
          <el-table-column sortable align="left" label="产品名称" prop="productName" width="120" />
          <el-table-column sortable align="left" label="产品sku" prop="productSku" width="120" />
          <el-table-column align="left" label="参考价" width="360" >
            <template #default="scope">
              <div v-if="scope.row.referPriceMin!==0&&scope.row.referPriceMax!==0">
                <div>
                  {{ (scope.row.referPriceMin).toFixed(2) }}~{{ (scope.row.referPriceMax).toFixed(2) }}$【美金】
                </div>
                <div>
                  {{ (scope.row.referPriceMin*scope.row.rate).toFixed(2) }}~{{ (scope.row.referPriceMax*scope.row.rate).toFixed(2) }}￥【填写汇率】{{scope.row.rate}}
                </div>
                <div>
                  {{ (scope.row.referPriceMin * rate.rate).toFixed(2) }}~{{ (scope.row.referPriceMax * rate.rate).toFixed(2) }}￥【实时汇率】{{ rate.rate.toFixed(2) }}
                </div>
              </div>
            </template>
          </el-table-column>
           <el-table-column label="参考图片" prop="referPics" width="200">
              <template #default="scope">
                 <div class="multiple-img-box">
                    <el-image preview-teleported v-for="(item,index) in scope.row.referPics" :key="index" style="width: 80px; height: 80px" :src="getUrl(item)" fit="cover"/>
                </div>
              </template>
           </el-table-column>
          <el-table-column sortable align="left" label="备注" prop="remark" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateProductFunc(scope.row)">编辑</el-button>
            <el-button  type="primary" link icon="delete" @click="deleteRow(scope.row)">{{scope.row.DeletedAt?'彻底删除':'删除'}}</el-button>
            <el-button  type="primary" link icon="edit" @click="restoreRow(scope.row)" v-if="scope.row.DeletedAt">恢复</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="品牌:"  prop="brandId" >
            <el-select  v-model="formData.brandId" placeholder="请选择品牌" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in dataSource.brandId" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
            <el-form-item label="产品名称:"  prop="productName" >
              <el-input v-model="formData.productName" :clearable="true"  placeholder="请输入产品名称" />
            </el-form-item>
            <el-form-item label="产品sku:"  prop="productSku" >
              <el-input v-model="formData.productSku" :clearable="true"  placeholder="请输入产品sku" />
            </el-form-item>
            <div>实时汇率:{{rate.rate}} 更新时间:{{formatDate(rate.time)}}</div>
            <el-form-item label="美元汇率:"  prop="rate" class="mt-2">
              <el-input-number v-model="formData.rate"  style="width:100%" :precision="2" :step="0.01" :clearable="true"  >
                <template #prefix>
                  <span class="invisible">美金/人民币</span>
                </template>
                <template #suffix>
                  <span>美金/人民币</span>
                </template>
              </el-input-number>
            </el-form-item>
            <el-form-item label="参考价（Min）:"  prop="referPriceMin" >
              <el-input-number v-model="formData.referPriceMin"  style="width:100%" :precision="2" :clearable="true" >
                <template #prefix>
                  <span class="invisible">美金</span>
                </template>
                <template #suffix>
                  <span>美金</span>
                </template>
              </el-input-number>
            </el-form-item>
            <el-form-item label="参考价（Max）:"  prop="referPriceMax" >
              <el-input-number v-model="formData.referPriceMax"  style="width:100%" :precision="2" :clearable="true"  >
                <template #prefix>
                  <span class="invisible">美金</span>
                </template>
                <template #suffix>
                  <span>美金</span>
                </template>
              </el-input-number>
            </el-form-item>
            <el-form-item label="参考图片:"  prop="referPics" >
                <SelectImage
                 multiple
                 v-model="formData.referPics"
                 file-type="image"
                 />
            </el-form-item>
            <el-form-item label="备注:"  prop="remark" >
              <el-input v-model="formData.remark" :clearable="true"  placeholder="请输入备注" />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="品牌">
                        {{ filterDataSource(dataSource.brandId,detailFrom.brandId) }}
                    </el-descriptions-item>
                    <el-descriptions-item label="产品名称">
                        {{ detailFrom.productName }}
                    </el-descriptions-item>
                    <el-descriptions-item label="产品sku">
                        {{ detailFrom.productSku }}
                    </el-descriptions-item>
                    <el-descriptions-item label="参考价">
                      <div v-if="detailFrom.referPriceMin&&detailFrom.referPriceMax&&detailFrom.referPriceMin!==0&&detailFrom.referPriceMax!==0">
                        <div>
                          {{ (detailFrom.referPriceMin).toFixed(2) }}~{{ (detailFrom.referPriceMax).toFixed(2) }}$【美金】
                        </div>
                        <div>
                          {{ (detailFrom.referPriceMin*detailFrom.rate).toFixed(2) }}~{{ (detailFrom.referPriceMax*detailFrom.rate).toFixed(2) }}￥【填写汇率】{{detailFrom.rate}}
                        </div>
                        <div>
                          {{ (detailFrom.referPriceMin * rate.rate).toFixed(2) }}~{{ (detailFrom.referPriceMax * rate.rate).toFixed(2) }}￥【实时汇率】{{ rate.rate.toFixed(2) }}
                        </div>
                      </div>
                    </el-descriptions-item>
                    <el-descriptions-item label="美元汇率">
                      {{detailFrom.rate}}
                    </el-descriptions-item>
                    <el-descriptions-item label="参考图片">
                            <el-image style="width: 50px; height: 50px; margin-right: 10px" :preview-src-list="returnArrImg(detailFrom.referPics)" :initial-index="index" v-for="(item,index) in detailFrom.referPics" :key="index" :src="getUrl(item)" fit="cover" />
                    </el-descriptions-item>
                    <el-descriptions-item label="备注">
                        {{ detailFrom.remark }}
                    </el-descriptions-item>
                    <el-descriptions-item label="创建时间">
                      {{ formatDate(detailFrom.CreatedAt) }}
                    </el-descriptions-item>
                    <el-descriptions-item label="更新时间">
                      {{ formatDate(detailFrom.UpdatedAt) }}
                    </el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
  import {
    getProductDataSource,
    createProduct,
    deleteProduct,
    deleteProductByIds,
    updateProduct,
    findProduct,
    getProductList, restoreProduct
  } from '@/api/bagique/product'
import { getRate } from '@/api/bagique/common'
import { getUrl } from '@/utils/image'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'


defineOptions({
    name: 'Product'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
            brandId: undefined,
            productName: '',
            productSku: '',
            referPriceMin: 0,
            referPriceMax: 0,
            rate:0,
            referPics: [],
            remark: '',
        })
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getProductDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()



// 验证规则
const rule = reactive({
               productName : [{
                   required: true,
                   message: '请正确填写产品名称',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               productSku : [{
                   required: true,
                   message: '请正确填写产品sku',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
              rate: [
                {
                  validator: (rule, value, callback) => {
                    //价格必须大于等于0
                    if (value < 0) {
                      callback(new Error('美元汇率不能小于0'))
                    } else {
                      callback()
                    }
                  },
                  trigger: ['input', 'blur']
                }
              ],
              referPriceMin: [
                {
                  validator: (rule, value, callback) => {
                    //价格必须大于等于0
                    if (value < 0) {
                      callback(new Error('参考价（Min）不能小于0'))
                    }
                    if (value !== undefined && formData.value.referPriceMax !== undefined) {
                      if (value > formData.value.referPriceMax) {
                        callback(new Error('参考价（Min）必须小于等于参考价（Max）'))
                      } else {
                        callback()
                      }
                    } else {
                      callback()
                    }
                  },
                  trigger: ['input', 'blur']
                }
              ],
              referPriceMax: [
                {
                  validator: (rule, value, callback) => {
                    //价格必须大于等于0
                    if (value < 0) {
                      callback(new Error('参考价（Max）不能小于0'))
                    }
                    if (formData.value.referPriceMin !== undefined && value !== undefined) {
                      if (value < formData.value.referPriceMin) {
                        callback(new Error('参考价（Max）必须大于等于参考价（Min）'))
                      } else {
                        callback()
                      }
                    } else {
                      callback()
                    }
                  },
                  trigger: ['input', 'blur']
                }
              ]
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const rate = ref({
  rate: 0,
  time: ''
})
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
            CreatedAt: 'created_at',
            brandId: 'brand_id',
            productName: 'product_name',
            productSku: 'product_sku',
            referPriceMin: 'refer_price_min',
            referPriceMax: 'refer_price_max',
            referPics: 'refer_pics',
            remark: 'remark',
  }

  let sort = sortMap[prop]
  if(!sort){
   sort = prop.replace(/[A-Z]/g, match => `_${match.toLowerCase()}`)
  }

  searchInfo.value.sort = sort
  searchInfo.value.order = order
  getTableData()
}

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getProductList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

// 获取美元汇率
const getRateData = async () => {
  const res = await getRate({})
  if (res.code === 0) {
    rate.value = res.data
    formData.value.rate = res.data.rate
  }
}

getTableData()
getRateData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  const tips = row.DeletedAt ? '确定要彻底删除吗?' : '确定要删除吗?'
    ElMessageBox.confirm(tips, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteProductFunc(row)
        })
    }

// 恢复行
const restoreRow = (row) => {
    ElMessageBox.confirm('确定要恢复吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            restoreProductFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteProductByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateProductFunc = async(row) => {
    const res = await findProduct({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteProductFunc = async (row) => {
    const res = await deleteProduct({ ID: row.ID,TYPE:row.DeletedAt?'HARD':'SOFT' })
    const tips = row.DeletedAt ? '彻底删除成功' : '删除成功'
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: tips
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 恢复行
const restoreProductFunc = async (row) => {
    const res = await restoreProduct({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '恢复成功'
            })
            getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
    getRateData()
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        brandId: undefined,
        productName: '',
        productSku: '',
        referPriceMin: 0,
        referPriceMax: 0,
        rate:0,
        referPics: [],
        remark: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}


const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findProduct({ ID: row.ID })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}


</script>

<style>

</style>