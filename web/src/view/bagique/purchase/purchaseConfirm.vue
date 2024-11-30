<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule"
               @keyup.enter="onSubmit">
        <el-form-item label="创建日期" prop="createdAt">
          <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
          </template>
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期"
                          :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期"
                          :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="searchInfo.remark" placeholder="搜索条件" />
        </el-form-item>
        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开
          </el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
        <div>实时汇率:{{ rate.rate }} 更新时间:{{ formatDate(rate.time) }}</div>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <ExportTemplate template-id="bagique_Evaluate" />
        <ExportExcel template-id="bagique_Evaluate" />
        <ImportExcel template-id="bagique_Evaluate" @on-success="getTableData" />
      </div>
      <el-table
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @sort-change="sortChange"
      >

        <el-table-column sortable align="left" label="创建日期" prop="CreatedAt" width="180" fixed="left">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>

        <el-table-column sortable align="left" label="产品" prop="productId" width="120" fixed="left">
          <template #default="scope">

            <span>{{ filterDataSource(dataSource.productId, scope.row.productId) }}</span>

          </template>
        </el-table-column>
        <el-table-column label="细节图" prop="evaluatePics" width="200">
          <template #default="scope">
            <div class="multiple-img-box">
              <el-image preview-teleported v-for="(item,index) in scope.row.evaluatePics" :key="index"
                        style="width: 80px; height: 80px" :src="getUrl(item)" fit="cover" />
            </div>
          </template>
        </el-table-column>
        <el-table-column label="公司估价" prop="evaluatePrices" align="center">
          <el-table-column label="【格式】估价/估价费用" align="center">
            <el-table-column v-for="(item,index) in dataSource.companyId" :key="index" :label="item.label"
                             align="center" width="300">
              <template #default="scope">
                <div v-if="scope.row.evaluatePrices.some((it) => it.companyId === item.value)">
                  <div>
                    <span>{{ (getEvaluatePrice(scope.row.evaluatePrices, item.value).price).toFixed(2)
                      }}$/{{ (getEvaluatePrice(scope.row.evaluatePrices, item.value).fee).toFixed(2) }}$【美金】</span>
                  </div>
                  <div>
                    <span>{{ (getEvaluatePrice(scope.row.evaluatePrices, item.value).price * scope.row.rate).toFixed(2)
                      }}￥/{{ (getEvaluatePrice(scope.row.evaluatePrices, item.value).fee * scope.row.rate).toFixed(2)
                      }}￥【填写汇率】{{ scope.row.rate }}</span>
                  </div>
                  <div>
                    <span>{{ (getEvaluatePrice(scope.row.evaluatePrices, item.value).price * rate.rate).toFixed(2)
                      }}￥/{{ (getEvaluatePrice(scope.row.evaluatePrices, item.value).fee * rate.rate).toFixed(2)
                      }}￥【实时汇率】{{ rate.rate }}</span>
                  </div>
                  <div>
                    <span>{{ getEvaluatePrice(scope.row.evaluatePrices, item.value).remark }}</span>【备注】
                  </div>
                </div>
                <div v-else>
                  - -
                </div>
              </template>
            </el-table-column>
          </el-table-column>
        </el-table-column>
        <el-table-column sortable align="left" label="备注" prop="remark" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
              <el-icon style="margin-right: 5px">
                <InfoFilled />
              </el-icon>
              查看
            </el-button>
            <el-button type="primary" link class="table-button" @click="updateEvaluateFunc(scope.row)">
              <el-icon style="margin-right: 5px">
                <InfoFilled />
              </el-icon>
              添加采购
            </el-button>
            <el-button type="primary" link class="table-button" @click="refuseRow(scope.row)">
              <el-icon style="margin-right: 5px">
                <InfoFilled />
              </el-icon>
              驳回采购
            </el-button>
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
          <span class="text-lg">添加采购</span>
          <div>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>

      <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="估价信息:" prop="evaluatePrices">
          <el-table ref="singleTableRef" :data="formData.evaluatePrices" highlight-current-row @current-change="handleCurrent">
            <el-table-column label="估价公司" align="center">
              <template #default="scope">
                <span>{{dataSource.companyId.find((item) => item.value === scope.row.companyId)?.label}}</span>
              </template>
            </el-table-column>
            <el-table-column label="估价[美金]" align="center">
              <template #default="scope">
                <span>{{scope.row.price.toFixed(2)}}</span>
              </template>
            </el-table-column>
            <el-table-column label="估价费用[美金]" align="center">
              <template #default="scope">
                <span>{{scope.row.fee.toFixed(2)}}</span>
              </template>
            </el-table-column>
            <el-table-column label="备注" align="center">
              <template #default="scope">
                <span>{{scope.row.remark}}</span>
              </template>
            </el-table-column>
            <el-table-column label="操作" align="center">
              <template #default="scope">
                <el-button :type="currentRow === scope.row ? 'danger':'primary'" @click="setCurrent(scope.row)">
                  {{currentRow === scope.row ? '已选' : '选择'}}
                </el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-form-item>
        <el-form-item label="备注:" prop="remark">
          <el-input v-model="formData.remark" :clearable="true" placeholder="请输入备注" />
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow"
               title="查看">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="产品">
          {{ filterDataSource(dataSource.productId, detailFrom.productId) }}
        </el-descriptions-item>
        <el-descriptions-item label="细节图">
          <el-image style="width: 50px; height: 50px; margin-right: 10px"
                    :preview-src-list="returnArrImg(detailFrom.evaluatePics)" :initial-index="index"
                    v-for="(item,index) in detailFrom.evaluatePics" :key="index" :src="getUrl(item)" fit="cover" />
        </el-descriptions-item>
        <el-descriptions-item label="公司估价">
          <div class="flex justify-center">
            【格式】估价/估价费用
          </div>
          <div class="flex flex-wrap justify-around">
            <div v-for="(item,index) in dataSource.companyId" :key="index" class="flex flex-col items-center">
              <div>
                <span>{{ filterDataSource(dataSource.companyId, item.value) }}</span>
              </div>
              <div v-if="detailFrom.evaluatePrices&&detailFrom.evaluatePrices.some((it) => it.companyId === item.value)"
                   class="flex flex-col items-center">
                <div>
                  <span>{{ (getEvaluatePrice(detailFrom.evaluatePrices, item.value).price).toFixed(2)
                    }}$/{{ (getEvaluatePrice(detailFrom.evaluatePrices, item.value).fee).toFixed(2) }}$【美金】</span>
                </div>
                <div>
                  <span>{{ (getEvaluatePrice(detailFrom.evaluatePrices, item.value).price * detailFrom.rate).toFixed(2)
                    }}￥/{{ (getEvaluatePrice(detailFrom.evaluatePrices, item.value).fee * detailFrom.rate).toFixed(2)
                    }}￥【填写汇率】{{ detailFrom.rate }}</span>
                </div>
                <div>
                  <span>{{ (getEvaluatePrice(detailFrom.evaluatePrices, item.value).price * rate.rate).toFixed(2)
                    }}￥/{{ (getEvaluatePrice(detailFrom.evaluatePrices, item.value).fee * rate.rate).toFixed(2)
                    }}￥【实时汇率】{{ rate.rate }}</span>
                </div>
                <div>
                  <span>{{ getEvaluatePrice(detailFrom.evaluatePrices, item.value).remark }}</span>【备注】
                </div>
              </div>
              <div v-else>
                - -
              </div>
            </div>
          </div>
        </el-descriptions-item>
        <el-descriptions-item label="美元汇率">
          {{ detailFrom.rate }}
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
    updateEvaluate,
    findEvaluate,
    getEvaluateList, restoreEvaluate
  } from '@/api/bagique/evaluate'
  import { getUrl } from '@/utils/image'
  // 图片选择组件

  // 全量引入格式化工具 请按需保留
  import {
    getDictFunc,
    formatDate,
    filterDataSource,
    returnArrImg
  } from '@/utils/format'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { ref, reactive } from 'vue'

  // 导出组件
  import ExportExcel from '@/components/exportExcel/exportExcel.vue'
  // 导入组件
  import ImportExcel from '@/components/exportExcel/importExcel.vue'
  // 导出模板组件
  import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'
  import { getRate } from '@/api/bagique/common'


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
    rate: 0
  })
  const dataSource = ref([])
  const getDataSourceFunc = async () => {
    const res = await getEvaluateDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()

  const currentRow = ref()
  const singleTableRef = ref()

  const setCurrent = (row) => {
    singleTableRef.value.setCurrentRow(row)
  }

  const handleCurrent = (val) => {
    currentRow.value = val
  }


  // 验证规则
  const rule = reactive({
    productId: [{
      required: true,
      message: '请正确选择产品',
      trigger: ['input', 'blur']
    }
    ],
    sellerId: [{
      required: true,
      message: '请正确选择卖家',
      trigger: ['input', 'blur']
    }
    ],
    evaluatePics: [{
      required: true,
      message: '请正确上传细节图',
      trigger: ['input', 'blur']
    }
    ],
    status: [{
      required: true,
      message: '请正确选择估价状态',
      trigger: ['input', 'blur']
    },
      {
        whitespace: true,
        message: '不能只输入空格',
        trigger: ['input', 'blur']
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
    evaluatePrices: [
      //companyId price fee
      //companyId不能为空 price和fee要大于等于0
      {
        validator: (rule, value, callback) => {
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
        }, required: true, trigger: ['input', 'blur']
      }
    ]
  })

  const searchRule = reactive({
    createdAt: [
      {
        validator: (rule, value, callback) => {
          if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
            callback(new Error('请填写结束日期'))
          } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
            callback(new Error('请填写开始日期'))
          } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
            callback(new Error('开始日期应当早于结束日期'))
          } else {
            callback()
          }
        }, trigger: 'change'
      }
    ]
  })

  const elFormRef = ref()
  const elSearchFormRef = ref()

  // =========== 表格控制部分 ===========
  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(10)
  const tableData = ref([])
  const searchInfo = ref({
    status: 'FINISH'
  })
  const rate = ref({
    rate: 0,
    time: ''
  })
  // 排序
  const sortChange = ({ prop, order }) => {
    const sortMap = {
      CreatedAt: 'created_at',
      productId: 'product_id',
      sellerId: 'seller_id',
      evaluatePics: 'evaluate_pics',
      status: 'status',
      remark: 'remark'
    }

    let sort = sortMap[prop]
    if (!sort) {
      sort = prop.replace(/[A-Z]/g, match => `_${match.toLowerCase()}`)
    }

    searchInfo.value.sort = sort
    searchInfo.value.order = order
    getTableData()
  }

  // 重置
  const onReset = () => {
    searchInfo.value = { status: 'FINISH' }
    getTableData()
  }

  // 搜索
  const onSubmit = () => {
    elSearchFormRef.value?.validate(async (valid) => {
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
  const getTableData = async () => {
    const table = await getEvaluateList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
  const setOptions = async () => {
    evaluate_statusOptions.value = await getDictFunc('evaluate_status')
  }

  // 获取需要的字典 可能为空 按需保留
  setOptions()

  // 行为控制标记（弹窗内部需要增还是改）
  const type = ref('')

  // 更新行
  const updateEvaluateFunc = async (row) => {
    const res = await findEvaluate({ ID: row.ID })
    if (res.code === 0) {
      formData.value = res.data
      dialogFormVisible.value = true
    }
  }


  // 弹窗控制标记
  const dialogFormVisible = ref(false)


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
      rate: 0
    }
  }
  // 弹窗确定
  const enterDialog = async () => {
    elFormRef.value?.validate(async (valid) => {
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
    if (!formData.value.evaluatePrices) {
      formData.value.evaluatePrices = []
    }
    formData.value.evaluatePrices.push({
      companyId: undefined,
      price: 0,
      fee: 0,
      remark: ''
    })
  }

  //缓存查询
  const getEvaluatePrice = (evaluatePrices, companyId) => {
    return evaluatePrices.find((item) => item.companyId === companyId)
  }

  //驳回估价
  const refuseEvaluateFunc = async (row) => {
    const res = await refuseEvaluate({ ID: row.ID })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: "驳回成功"
      })
      if (tableData.value.length === 1 && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  }

  //确定驳回
  const refuseRow = (row) => {
    ElMessageBox.confirm('确定要驳回估价吗?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      refuseEvaluateFunc(row)
    })
  }

</script>

<style>

</style>