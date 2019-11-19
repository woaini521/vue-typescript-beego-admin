<template>
  <div class="app-container">
    <cmp-edit ref="edit" :show.sync="showEdit" @edit="getData"></cmp-edit>
    <!---/ 按钮组 -->
    <el-row class="table-btn-row">
      <el-button @click="onAdd" size="medium" type="primary">添加</el-button>
    </el-row>
    <!---\ 按钮组 -->
    <el-table v-loading="listLoading" :data="list" element-loading-text="Loading" border fit highlight-current-row>
      <el-table-column align="center" label="ID" width="95">
        <template slot-scope="scope">{{ scope.row.id }}</template>
      </el-table-column>
      <el-table-column label="登录名">
        <template slot-scope="scope">{{ scope.row.loginName }}</template>
      </el-table-column>
      <el-table-column label="昵称" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.realName }}</span>
        </template>
      </el-table-column>
      <el-table-column label="手机号" width="120" align="center">
        <template slot-scope="scope">{{ scope.row.phone }}</template>
      </el-table-column>
      <el-table-column label="邮件" align="center">
        <template slot-scope="scope">{{ scope.row.email }}</template>
      </el-table-column>
      <el-table-column label="上次登录" width="160" align="center">
        <template slot-scope="scope">{{ scope.row.lastLogin }}</template>
      </el-table-column>
      <el-table-column label="登录IP" width="160" align="center">
        <template slot-scope="scope">{{ scope.row.lastIP }}</template>
      </el-table-column>
      <el-table-column label="创建时间" width="160" align="center">
        <template slot-scope="scope">{{ scope.row.createTime }}</template>
      </el-table-column>
      <el-table-column label="更新时间" width="160" align="center">
        <template slot-scope="scope">{{ scope.row.updateTime }}</template>
      </el-table-column>
      <el-table-column class-name="status-col" label="状态" align="center">
        <template slot-scope="scope">
          <el-tag :type="scope.row.status | statusTypeFilter">{{ scope.row.status | statusDescFilter }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" label="操作栏" width="250">
        <template slot-scope="scope">
          <el-button @click="handleTools(scope.row, 'edit')" size="mini" type="primary">编辑</el-button>
          <el-button @click="handleTools(scope.row, 'status')" size="mini" type="warning">启|禁</el-button>
          <el-button @click="handleTools(scope.row, 'del')" size="mini" type="danger">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <!-- <el-pagination
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :current-page="listQuery.page"
      :page-sizes="paginationConfig.pageSizes"
      :page-size="paginationConfig.pageSize"
      :layout="paginationConfig.layout"
      :total="total"
    ></el-pagination> -->
    <cmp-pagination
      v-show="total > 0"
      :total="total"
      :page.sync="listQuery.page"
      :limit.sync="listQuery.limit"
      @pagination="getData"
    />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { getUsers, editUser, delAdmin } from '@/api/users'
import { IAdminUserData } from '@/api/types'
import { TableConfig } from '@/utils/pageConfig'
import CmpPagination from '@/components/Pagination/index.vue'
import CmpEdit from './edit.vue'
import { AdminEditDefaultForm } from './types'

@Component({
  name: 'AdminUser',
  filters: {
    statusTypeFilter: (status: number) => {
      const statusMap: { [key: number]: string } = {
        0: 'danger',
        1: 'success'
      }
      return statusMap[status]
    },
    statusDescFilter: (status: number) => {
      const statusMap: { [key: number]: string } = {
        0: '禁用',
        1: '正常'
      }
      return statusMap[status]
    },
    parseTime: (timestamp: string) => {
      return new Date(timestamp).toISOString()
    }
  },
  components: {
    CmpEdit,
    CmpPagination
  }
})
export default class extends Vue {
  public showEdit = false

  private list: IAdminUserData[] = []
  private total = 0

  private listLoading = true
  private listQuery = {
    page: 1,
    limit: 20
  }
  private paginationConfig = TableConfig

  created() {
    this.getData()
  }

  private handleTools(data: IAdminUserData, type: string) {
    if (data.id === 1) {
      this.$message.error('超级管理员不允许操作')
      return
    }
    switch (type) {
      case 'edit':
        const roleIds: number[] = data.roleIds.split(',').map(v => parseInt(v + ''))
        ;(this.$refs as any).edit.setForm({
          id: data.id,
          loginName: data.loginName,
          realName: data.realName,
          phone: data.phone,
          email: data.email,
          roleIds: roleIds
        })
        this.showEdit = true
        break
      case 'status':
        this.$confirm(`是否真的${data.status === 1 ? '禁用' : '启用'}【${data.realName}】？`).then(() =>
          editUser({
            id: data.id,
            status: data.status === 0 ? 1 : 0,
            editType: 'STATUS'
          }).then(() => this.getData())
        )
        break
      case 'del':
        this.$confirm(`是否真的删除【${data.realName}】？`).then(() => {
          delAdmin(data.id).then(() => this.getData())
        })
        break
      default:
        this.$message.warning('unknown event!')
    }
  }

  private onAdd() {
    this.showEdit = true
    ;(this.$refs as any).edit.setForm(AdminEditDefaultForm)
  }

  private async getData() {
    try {
      this.listLoading = true
      const { data } = await getUsers(this.listQuery)
      const { list = [], total = 0 } = data
      this.list = list
      this.total = total
    } catch (error) {}

    this.listLoading = false
  }
}
</script>
