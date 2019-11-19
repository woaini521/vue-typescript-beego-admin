<template>
  <div class="app-container">
    <cmp-edit ref="edit" :show.sync="showEdit" @edit="getData" />

    <!---/ 按钮组 -->
    <el-row class="table-btn-row">
      <el-button @click="onAdd" size="medium" type="primary">添加</el-button>
    </el-row>
    <!---\ 按钮组 -->
    <el-table v-loading="listLoading" :data="list" element-loading-text="Loading" border fit highlight-current-row>
      <el-table-column align="center" label="ID" width="95">
        <template slot-scope="scope">{{ scope.row.id }}</template>
      </el-table-column>
      <el-table-column label="权限名称">
        <template slot-scope="scope">{{ scope.row.roleName }}</template>
      </el-table-column>
      <el-table-column label="创建时间" width="160" align="center">
        <template slot-scope="scope">{{ scope.row.createTime }}</template>
      </el-table-column>
      <el-table-column label="更新时间" width="160" align="center">
        <template slot-scope="scope">{{ scope.row.updateTime }}</template>
      </el-table-column>
      <el-table-column label="备注" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.detail }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="操作栏" width="250">
        <template slot-scope="scope">
          <el-button @click="handlerTools(scope.row, 'edit')" size="mini" type="primary">编辑</el-button>
          <el-button @click="handlerTools(scope.row, 'del')" size="mini" type="danger">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
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
import { getRoles, getRoleAuth, delRole } from '@/api/role'
import { IAdminRole } from '@/api/types'

import { isArray } from '@/utils/validate'

import CmpPagination from '@/components/Pagination/index.vue'
import CmpEdit from './edit.vue'

import { RoleEditDefaultForm } from './types'

@Component({
  name: 'AdminUser',
  components: {
    CmpEdit,
    CmpPagination
  }
})
export default class extends Vue {
  public showEdit = false
  private list: IAdminRole[] = []
  private listLoading = true

  public total = 0
  public listQuery = {
    page: 1,
    limit: 20
  }

  created() {
    this.getData()
  }

  private handlerTools(data: IAdminRole, type: string) {
    switch (type) {
      case 'edit':
        this.showEdit = true
        getRoleAuth(data.id)
          .then(res => {
            console.error('auths', res)

            const { authIds = [] } = res.data
            isArray(authIds) && (data.authIds = authIds)
            console.error('auths', authIds)
            ;(this.$refs as any).edit.setForm(data)
          })
          .catch(() => {
            ;(this.$refs as any).edit.setForm(data)
          })

        break
      case 'del':
        this.$confirm(`是否真的删除【${data.roleName}】？`).then(() => {
          delRole(data.id).then(() => this.getData())
        })
        break
      default:
        this.$message.warning('unknown event!')
    }
  }

  private onAdd() {
    this.showEdit = true
    ;(this.$refs as any).edit.setForm(RoleEditDefaultForm)
  }

  public async getData() {
    try {
      this.listLoading = true
      const { data } = await getRoles(this.listQuery)
      const { list = [], total = 0 } = data
      this.list = list
      this.total = total
    } catch (error) {}

    this.listLoading = false
  }
}
</script>
