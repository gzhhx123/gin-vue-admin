
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
      
            <el-form-item label="产品" prop="productId">
               <el-select  v-model="searchInfo.productId" placeholder="请选择产品" :clearable="true" >
                   <el-option v-for="(item,key) in dataSource.productId" :key="key" :label="item.label" :value="item.value" />
               </el-select>
            </el-form-item>
            <el-form-item label="卖家" prop="sellerId">
               <el-select  v-model="searchInfo.sellerId" placeholder="请选择卖家" :clearable="true" >
                   <el-option v-for="(item,key) in dataSource.sellerId" :key="key" :label="item.label" :value="item.value" />
               </el-select>
            </el-form-item>
           <el-form-item label="估价状态" prop="status">
            <el-select v-model="searchInfo.status" clearable placeholder="请选择" @clear="()=>{searchInfo.status=undefined}">
              <el-option v-for="(item,key) in evaluate_statusOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
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
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            <ExportTemplate  template-id="bagique_Evaluate" />
            <ExportExcel  template-id="bagique_Evaluate" />
            <ImportExcel  template-id="bagique_Evaluate" @on-success="getTableData" />
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
        
        <el-table-column sortable align="left" label="产品" prop="productId" width="120">
          <template #default="scope">
                
                    <span>{{ filterDataSource(dataSource.productId,scope.row.productId) }}</span>
                
         </template>
         </el-table-column>
        <el-table-column sortable align="left" label="卖家" prop="sellerId" width="120">
          <template #default="scope">
                
                    <span>{{ filterDataSource(dataSource.sellerId,scope.row.sellerId) }}</span>
                
         </template>
         </el-table-column>
           <el-table-column label="细节图" prop="evaluatePics" width="200">
              <template #default="scope">
                 <div class="multiple-img-box">
                    <el-image preview-teleported v-for="(item,index) in scope.row.evaluatePics" :key="index" style="width: 80px; height: 80px" :src="getUrl(item)" fit="cover"/>
                </div>
              </template>
           </el-table-column>
        <el-table-column sortable align="left" label="估价状态" prop="status" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.status,evaluate_statusOptions) }}
            </template>
        </el-table-column>
          <el-table-column sortable align="left" label="备注" prop="remark" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateEvaluateFunc(scope.row)">编辑</el-button>
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
            <el-form-item label="产品:"  prop="productId" >
            <el-select  v-model="formData.productId" placeholder="请选择产品" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in dataSource.productId" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
            <el-form-item label="卖家:"  prop="sellerId" >
            <el-select  v-model="formData.sellerId" placeholder="请选择卖家" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in dataSource.sellerId" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
            <el-form-item label="细节图:"  prop="evaluatePics" >
                <SelectImage
                 multiple
                 v-model="formData.evaluatePics"
                 file-type="image"
                 />
            </el-form-item>
            <el-form-item label="估价状态:"  prop="status" >
              <el-select v-model="formData.status" placeholder="请选择估价状态" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in evaluate_statusOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="估价信息:" prop="evaluatePrices">
              <el-button type="primary" icon="edit" @click="addEvaluatePrice">
                新增估价信息
              </el-button>
              <el-table :data="formData.evaluatePrices">
                <el-table-column label="估价公司" align="center">
                  <template #default="scope">
                    <el-select
                      v-model="scope.row.companyId"
                      placeholder="请选择估价公司"
                      :clearable="true">
                      <el-option
                        v-for="(item,key) in dataSource.companyId"
                        :key="key"
                        :disabled="formData.evaluatePrices.some((it) => it.companyId === item.value)"
                        :label="item.label"
                        :value="item.value"/>
                    </el-select>
                  </template>
                </el-table-column>
                <el-table-column label="估价费用" align="center">
                  <template #default="scope">
                    <el-input-number v-model="scope.row.fee" placeholder="请输入估价费用" :precision="2" :clearable="true"  />
                  </template>
                </el-table-column>
                <el-table-column label="估价" align="center">
                  <template #default="scope">
                    <el-input-number v-model="scope.row.price" placeholder="请输入估价" :precision="2" :clearable="true"  />
                  </template>
                </el-table-column>
                <el-table-column align="center">
                  <template #default="scope">
                    <el-button type="danger" icon="delete" @click="() => formData.evaluatePrices.splice(scope.$index, 1)">
                      删除
                    </el-button>
                  </template>
                </el-table-column>
              </el-table>
            </el-form-item>
            <el-form-item label="备注:"  prop="remark" >
              <el-input v-model="formData.remark" :clearable="true"  placeholder="请输入备注" />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="产品">
                      {{filterDataSource(dataSource.productId,detailFrom.productId)}}
                    </el-descriptions-item>
                    <el-descriptions-item label="卖家">
                      {{filterDataSource(dataSource.sellerId,detailFrom.sellerId)}}
                    </el-descriptions-item>
                    <el-descriptions-item label="细节图">
                            <el-image style="width: 50px; height: 50px; margin-right: 10px" :preview-src-list="returnArrImg(detailFrom.evaluatePics)" :initial-index="index" v-for="(item,index) in detailFrom.evaluatePics" :key="index" :src="getUrl(item)" fit="cover" />
                    </el-descriptions-item>
                    <el-descriptions-item label="估价状态">
                      {{ filterDict(detailFrom.status,evaluate_statusOptions) }}
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
    getEvaluateDataSource,
    createEvaluate,
    deleteEvaluate,
    deleteEvaluateByIds,
    updateEvaluate,
    findEvaluate,
    getEvaluateList, restoreEvaluate
  } from '@/api/bagique/evaluate'
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
    name: 'Evaluate'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const evaluate_statusOptions = ref([])
const formData = ref({
            productId: undefined,
            sellerId: undefined,
            evaluatePrices: [],
            evaluatePics: [],
            status: '',
            remark: '',
        })
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getEvaluateDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()



// 验证规则
const rule = reactive({
               productId : [{
                   required: true,
                   message: '请正确选择产品',
                   trigger: ['input','blur'],
               },
              ],
               sellerId : [{
                   required: true,
                   message: '请正确选择卖家',
                   trigger: ['input','blur'],
               },
              ],
               evaluatePics : [{
                   required: true,
                   message: '请正确上传细节图',
                   trigger: ['input','blur'],
               },
              ],
               status : [{
                   required: true,
                   message: '请正确选择估价状态',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
              evaluatePrices: [
                //companyId price fee
                //companyId不能为空 price和fee要大于等于0
                { validator: (rule, value, callback) => {
                  if (value.length === 0) {
                    callback(new Error('请至少添加一条估价信息'))
                  } else {
                    for (let i = 0; i < value.length; i++) {
                      if (!value[i].companyId) {
                        callback(new Error('请正确选择估价公司'))
                        return
                      }
                      if (value[i].price < 0) {
                        callback(new Error('估价不能小于0'))
                        return
                      }
                      if (value[i].fee < 0) {
                        callback(new Error('估价费用不能小于0'))
                        return
                      }
                    }
                    callback()
                  }
                }, required: true,trigger: ['input', 'blur'] }
              ],
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
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
            CreatedAt: 'created_at',
            productId: 'product_id',
            sellerId: 'seller_id',
            evaluatePics: 'evaluate_pics',
            status: 'status',
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
  const table = await getEvaluateList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    evaluate_statusOptions.value = await getDictFunc('evaluate_status')
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
    deleteEvaluateFunc(row)
  })
}

// 恢复行
const restoreRow = (row) => {
  ElMessageBox.confirm('确定要恢复吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    restoreEvaluateFunc(row)
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
      const res = await deleteEvaluateByIds({ IDs })
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
const updateEvaluateFunc = async(row) => {
    const res = await findEvaluate({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteEvaluateFunc = async (row) => {
  const res = await deleteEvaluate({ ID: row.ID, TYPE: row.DeletedAt ? 'HARD' : 'SOFT' })
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
const restoreEvaluateFunc = async (row) => {
  const res = await restoreEvaluate({ ID: row.ID })
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
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        productId: undefined,
        sellerId: undefined,
        evaluatePrices: [],
        evaluatePics: [],
        status: '',
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
  const res = await findEvaluate({ ID: row.ID })
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

// 新增估价信息
const addEvaluatePrice = () => {
  if(!formData.value.evaluatePrices){
      formData.value.evaluatePrices = []
  }
  formData.value.evaluatePrices.push({
    companyId: undefined,
    price: 0,
    fee: 0,
  })
}


</script>

<style>

</style>