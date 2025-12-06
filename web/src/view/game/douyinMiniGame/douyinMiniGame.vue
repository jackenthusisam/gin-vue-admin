<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="游戏名称">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="AppID">
          <el-input v-model="searchInfo.appId" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDeleteIds">删除</el-button>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="ID" prop="ID" width="80" />
        <el-table-column align="left" label="游戏名称" prop="name" width="120" />
        <el-table-column align="left" label="AppID" prop="appId" width="120" />
        <el-table-column align="left" label="AppSecret" prop="appSecret" width="120" />
        <el-table-column align="left" label="状态" prop="status" width="120">
          <template #default="scope">
            <el-tag :type="scope.row.status === 1 ? 'success' : 'info'">{{ scope.row.status === 1 ? '开启' : '关闭' }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="描述" prop="description" width="200" />
        <el-table-column align="left" label="操作">
          <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateDouyinMiniGameFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type==='create'?'添加':'修改'" destroy-on-close>
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="游戏名称" prop="name">
          <el-input v-model="formData.name" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="AppID" prop="appId">
          <el-input v-model="formData.appId" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="AppSecret" prop="appSecret">
          <el-input v-model="formData.appSecret" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="formData.status" placeholder="请选择">
            <el-option label="开启" :value="1" />
            <el-option label="关闭" :value="2" />
          </el-select>
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="formData.description" type="textarea" placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createDouyinMiniGame,
  deleteDouyinMiniGame,
  deleteDouyinMiniGameByIds,
  updateDouyinMiniGame,
  findDouyinMiniGame,
  getDouyinMiniGameList
} from '@/api/game/douyinMiniGame'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'DouyinMiniGame'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  name: '',
  appId: '',
  appSecret: '',
  status: 1,
  description: '',
})

// 验证规则
const rule = reactive({
  name: [{ required: true, message: '请输入游戏名称', trigger: 'blur' }],
  appId: [{ required: true, message: '请输入AppID', trigger: 'blur' }],
  status: [{ required: true, message: '请选择状态', trigger: 'change' }],
})

const searchInfo = ref({})
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const multipleSelection = ref([])

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
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
  const table = await getDouyinMiniGameList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// 批量操作
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

const onDeleteIds = async() => {
  const ids = []
  multipleSelection.value.forEach(item => {
    ids.push(item.ID)
  })
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const res = await deleteDouyinMiniGameByIds({ ids })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === ids.length && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')
const dialogFormVisible = ref(false)

// 更新行
const updateDouyinMiniGameFunc = async(row) => {
  const res = await findDouyinMiniGame({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reDouyinMiniGame
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const res = await deleteDouyinMiniGame(row)
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
  })
}

// 弹窗控制标记
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    name: '',
    appId: '',
    appSecret: '',
    status: 1,
    description: '',
  }
}

// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createDouyinMiniGame(formData.value)
        break
      case 'update':
        res = await updateDouyinMiniGame(formData.value)
        break
      default:
        res = await createDouyinMiniGame(formData.value)
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

const elFormRef = ref(null)
</script>

<style>
</style>
