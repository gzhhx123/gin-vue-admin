
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
        <div>实时汇率:{{ rate.rate }} 更新时间:{{ formatDate(rate.time) }}</div>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
<!--            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>-->
            <ExportTemplate  template-id="bagique_Purchase" />
            <ExportExcel  template-id="bagique_Purchase" />
            <ImportExcel  template-id="bagique_Purchase" @on-success="getTableData" />
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
<!--        <el-table-column type="selection" width="55" />-->
        
        <el-table-column align="left" label="创建日期" prop="CreatedAt" width="180" fixed="left">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column sortable align="left" label="产品" prop="productId" width="120" fixed="left">
          <template #default="scope">
            {{ filterDataSource(dataSource.productId,scope.row.evaluate.productId) }}
          </template>
        </el-table-column>
        <el-table-column label="细节图" prop="evaluatePics" width="200">
          <template #default="scope">
            <div class="multiple-img-box">
              <el-image preview-teleported v-for="(item,index) in scope.row.evaluate.evaluatePics" :key="index" style="width: 80px; height: 80px" :src="getUrl(item)" fit="cover"/>
            </div>
          </template>
        </el-table-column>
        <el-table-column sortable align="left" label="采购状态" prop="status" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.status,purchase_statusOptions) }}
          </template>
        </el-table-column>
        <el-table-column label="公司估价" prop="evaluatePrice" align="center">
          <el-table-column label="公司" align="center">
              <template #default="scope">
                {{ filterDataSource(dataSource.companyId,scope.row.evaluatePrice.companyId) }}
              </template>
          </el-table-column>
          <el-table-column label="估价" align="center" width="300">
            <template #default="scope">
              <div>
                {{ scope.row.evaluatePrice.price.toFixed(2) }}$【美金】
              </div>
              <div>
                {{(scope.row.evaluate.rate * scope.row.evaluatePrice.price).toFixed(2)}}￥【填写汇率】{{scope.row.evaluate.rate}}
              </div>
              <div>
                {{(rate.rate * scope.row.evaluatePrice.price).toFixed(2)}}￥【实时汇率】{{ rate.rate }}
              </div>
            </template>
          </el-table-column>
          <el-table-column label="估价费用" align="center">
            <template #default="scope">
              <div>
                {{ scope.row.evaluatePrice.fee.toFixed(2) }}$【美金】
              </div>
              <div>
                {{(scope.row.evaluate.rate * scope.row.evaluatePrice.fee).toFixed(2)}}￥【填写汇率】{{scope.row.evaluate.rate}}
              </div>
              <div>
                {{(rate.rate * scope.row.evaluatePrice.fee).toFixed(2)}}￥【实时汇率】{{ rate.rate }}
              </div>
            </template>
          </el-table-column>
          <el-table-column label="备注" align="center">
            <template #default="scope">
              {{ scope.row.evaluatePrice.remark }}
            </template>
          </el-table-column>
        </el-table-column>
        <el-table-column sortable align="left" label="备注" prop="remark" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="info-filled" class="table-button" @click="purchaseTimeAxisFunc(scope.row)">查看时间轴</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" v-if="scope.row.status==='ADD'" @click="updateTrackNoFunc(scope.row)">编辑物流信息</el-button>
              <el-button  type="primary" v-if="scope.row.status==='ADD'" link icon="circle-check" class="table-button" @click="finishRow(scope.row)">完成采购</el-button>
              <el-button  type="primary" v-if="scope.row.status==='ADD'||scope.row.status==='REFUSE'" link icon="circle-close" class="table-button" @click="cancelRow(scope.row)">取消采购</el-button>
            <el-button  v-if="!scope.row.DeletedAt&&scope.row.status!=='FINISH'" type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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
    <el-drawer destroy-on-close size="1600" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
         <template #header>
            <div class="flex justify-between items-center">
              <span class="text-lg">{{type==='trackNo'?'编辑物流信息':'编辑时间轴'}}</span>
              <div>
                <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                <el-button @click="closeDialog">取 消</el-button>
              </div>
            </div>
          </template>

          <el-form  label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="物流信息:">
              <el-button type="primary" icon="edit" @click="addTrackNo">
                新增物流信息
              </el-button>
              <el-table ref="singleTableRef" :data="trackData.list" >
                <el-table-column label="物流公司" align="center">
                  <template #default="scope">
                    <el-select
                      style="width: 100%"
                      v-model="scope.row.trackCompanyId"
                      placeholder="请选择物流公司"
                      :clearable="true">
                      <el-option
                        v-for="(item,key) in dataSource.trackCompanyId"
                        :key="key"
                        :label="item.label"
                        :value="item.value"/>
                    </el-select>
                  </template>
                </el-table-column>
                <el-table-column label="物流单号" align="center">
                  <template #default="scope">
                    <el-input
                      v-model="scope.row.orderSn"
                      placeholder="请输入物流单号"
                      clearable/>
                  </template>
                </el-table-column>
                <el-table-column label="单号类型" align="center">
                  <template #default="scope">
                    <el-select
                      style="width: 100%"
                      v-model="scope.row.type"
                      placeholder="请选择单号类型"
                      :clearable="true">
                      <el-option
                        v-for="(item,key) in track_no_typeOptions"
                        :key="key"
                        :label="item.label"
                        :value="item.value"/>
                    </el-select>
                  </template>
                </el-table-column>
                <el-table-column label="状态" align="center">
                  <template #default="scope">
                    <el-select
                      style="width: 100%"
                      v-model="scope.row.status"
                      placeholder="请选择状态"
                      :clearable="true">
                      <el-option
                        v-for="(item,key) in track_no_statusOptions"
                        :key="key"
                        :label="item.label"
                        :value="item.value"/>
                    </el-select>
                  </template>
                </el-table-column>
                <el-table-column label="备注" align="center">
                  <template #default="scope">
                    <el-input
                      v-model="scope.row.remark"
                      placeholder="请输入备注"
                      clearable/>
                  </template>
                </el-table-column>
                <el-table-column label="创建时间" align="center">
                  <template #default="scope">
                    {{ formatDate(scope.row.CreatedAt) }}
                  </template>
                </el-table-column>
                <el-table-column label="更新时间" align="center">
                  <template #default="scope">
                    {{ formatDate(scope.row.UpdatedAt) }}
                  </template>
                </el-table-column>
                <el-table-column label="操作" align="center">
                  <template #default="scope">
                    <div class="flex flex-col items-center justify-center">
                      <el-button @click="upRow(scope.row)">上移</el-button>
                      <el-button class="mt-2 ml-0"  @click="downRow(scope.row)">下移</el-button>
                      <el-button type="primary" class="mt-2 ml-0"  @click="removeRow(scope.row)">删除</el-button>
                    </div>
                  </template>
                </el-table-column>
              </el-table>
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="产品" v-if="Object.keys(detailFrom).length>0">
                        {{ filterDataSource(dataSource.productId,detailFrom.evaluate.productId) }}
                    </el-descriptions-item>
                    <el-descriptions-item label="细节图" v-if="Object.keys(detailFrom).length>0">
                      <el-image style="width: 50px; height: 50px; margin-right: 10px" :preview-src-list="returnArrImg(detailFrom.evaluate.evaluatePics)" :initial-index="index" v-for="(item,index) in detailFrom.evaluate.evaluatePics" :key="index" :src="getUrl(item)" fit="cover" />
                    </el-descriptions-item>
                    <el-descriptions-item label="估价公司估价" v-if="Object.keys(detailFrom).length>0">
                        <div>
                          <div>公司:{{ filterDataSource(dataSource.companyId,detailFrom.evaluatePrice.companyId) }}</div>
                          <br>
                          <div>
                            估价:<br>
                            {{ detailFrom.evaluatePrice.price.toFixed(2) }}$【美金】<br>
                            {{(detailFrom.evaluate.rate * detailFrom.evaluatePrice.price).toFixed(2)}}￥【填写汇率】{{detailFrom.evaluate.rate}}<br>
                            {{(rate.rate * detailFrom.evaluatePrice.price).toFixed(2)}}￥【实时汇率】{{ rate.rate }}
                          </div>
                          <br>
                          <div>
                            估价费用:<br>
                            {{ detailFrom.evaluatePrice.fee.toFixed(2) }}$【美金】<br>
                            {{(detailFrom.evaluate.rate * detailFrom.evaluatePrice.fee).toFixed(2)}}￥【填写汇率】{{detailFrom.evaluate.rate}}<br>
                            {{(rate.rate * detailFrom.evaluatePrice.fee).toFixed(2)}}￥【实时汇率】{{ rate.rate }}
                          </div>
                          <br>
                          <div>备注:{{ detailFrom.evaluatePrice.remark }}</div>
                        </div>
                    </el-descriptions-item>
                    <el-descriptions-item label="采购状态">
                        {{ filterDict(detailFrom.status,purchase_statusOptions) }}
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
    <el-drawer destroy-on-close size="800" v-model="axisShow" :show-close="true" :before-close="closeAxisShow" title="查看时间轴">
      <div>
        <el-steps direction="vertical" :active="timeAxis.length" class="w-full"  style="height: 40vh">
          <el-step :title="item['content']" :description="'达成时间:'+formatDate(item['time'])" v-for="(item,index) in timeAxis" />
        </el-steps>
      </div>
    </el-drawer>
  </div>
</template>

<script setup>
  import {
    createPurchase,
    deletePurchase,
    deletePurchaseByIds,
    updatePurchase,
    findPurchase,
    getPurchaseList,
    getPurchaseDataSource,
    findPurchaseTrackNosById,
    updateTrackNo,
    purchaseTimeAxis,
    cancelPurchase,
    finishPurchase
  } from '@/api/bagique/purchase'

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
import { getUrl } from '@/utils/image'
import { getProductDataSource } from '@/api/bagique/product'
  import { getRate } from '@/api/bagique/common'


defineOptions({
    name: 'Purchase'
})

// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const purchase_statusOptions = ref([])
const track_no_statusOptions = ref([])
const track_no_typeOptions = ref([])
const trackData = ref({list:[],id:0})
const dataSource = ref([])
const getDataSourceFunc = async()=>{
  const res = await getPurchaseDataSource()
  if (res.code === 0) {
    dataSource.value = res.data
  }
}
getDataSourceFunc()


// 验证规则
const rule = reactive({
               evaluatePriceId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               status : [{
                   required: true,
                   message: '请正确选择采购状态',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
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
const rate = ref({
  rate: 0,
  time: ''
})
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
            evaluatePriceId: 'evaluate_price_id',
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
  const table = await getPurchaseList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    }
  }

getTableData()
getRateData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    purchase_statusOptions.value = await getDictFunc('purchase_status')
    track_no_statusOptions.value = await getDictFunc('track_no_status')
    track_no_typeOptions.value = await getDictFunc('track_no_type')
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
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deletePurchaseFunc(row)
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
      const res = await deletePurchaseByIds({ IDs })
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
const updatePurchaseFunc = async(row) => {
    const res = await findPurchase({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}

// 更新行物流单号
  const updateTrackNoFunc = async(row) =>{
    const res = await findPurchaseTrackNosById({ ID: row.ID })
    type.value = 'trackNo'
    if (res.code === 0) {
      trackData.value.list = res.data.list
      trackData.value.id = row.ID
      dialogFormVisible.value = true
    }
  }


// 删除行
const deletePurchaseFunc = async (row) => {
    const res = await deletePurchase({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)


// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    trackData.value = {
      list:[],
      id:0
    }
}
// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     // elFormRef.value?.validate( async (valid) => {
     //         if (!valid) return btnLoading.value = false
     //
     //  })
  let res
  switch (type.value) {
    case 'trackNo':
      //对sort进行重新赋值
      trackData.value.list.forEach((item,index)=>{
        item.sort = index
      })
      res = await updateTrackNo(trackData.value)
      break
    case 'update':
      res = await updatePurchase(formData.value)
      break
    default:
      res = await createPurchase(formData.value)
      break
  }
  btnLoading.value = false
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '更新成功'
    })
    closeDialog()
    getTableData()
  }
}



const detailFrom = ref({})
  const timeAxis = ref([])

// 查看详情控制标记
const detailShow = ref(false)
  const axisShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findPurchase({ ID: row.ID })
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

const closeAxisShow = () => {
  axisShow.value = false
  timeAxis.value = []
}

// 新增物流信息
const addTrackNo = () => {
  if (!trackData.value.list){
    trackData.value.list = []
  }
  console.log(detailFrom.value)
  trackData.value.list.push({
    trackCompanyId: undefined,
    orderSn: '',
    status: '',
    type:'',
    remark: '',
    sort: trackData.value.list.length
  })
}

// 上移
const upRow = (row) => {
  const index = trackData.value.list.indexOf(row)
  if (index > 0) {
    const temp = trackData.value.list[index]
    trackData.value.list[index] = trackData.value.list[index - 1]
    trackData.value.list[index - 1] = temp
  }
}

// 下移
const downRow = (row) => {
  const index = trackData.value.list.indexOf(row)
  if (index < trackData.value.list.length - 1) {
    const temp = trackData.value.list[index]
    trackData.value.list[index] = trackData.value.list[index + 1]
    trackData.value.list[index + 1] = temp
  }
}

// 删除
const removeRow = (row) => {
  const index = trackData.value.list.indexOf(row)
  trackData.value.list.splice(index, 1)
}

//查看采购时间轴
  const purchaseTimeAxisFunc = async (row) => {
    const res = await purchaseTimeAxis({ ID: row.ID })
    if (res.code === 0) {
      timeAxis.value = res.data.list
      axisShow.value = true
    }
  }

  //完成采购
  const finishRow = (row) => {
    ElMessageBox.confirm('确定要完成采购吗?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      finishPurchaseFunc(row)
    })
  }

  // 取消采购
  const cancelRow = (row) => {
    ElMessageBox.confirm('确定要取消采购吗?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      cancelPurchaseFunc(row)
    })
  }


  // 完成行
  const finishPurchaseFunc = async (row)=>{
    const res = await finishPurchase({ ID: row.ID })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '完成采购成功'
      })
      getTableData()
    }
  }

  // 取消行
  const cancelPurchaseFunc = async (row)=>{
    const res = await cancelPurchase({ ID: row.ID })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '取消采购成功'
      })
      getTableData()
    }
  }


</script>

<style>

</style>
