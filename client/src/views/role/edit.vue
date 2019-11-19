<template>
  <el-dialog
    :title="`角色${form.id > 0 ? '编辑' : '添加'}`"
    :visible.sync="dialog"
    @close="handleDialogClose"
    size="100%"
  >
    <el-form ref="form" :model="form" :rules="formRule" label-width="80px" style="width: 400px; margin-left:50px;">
      <el-form-item label="权限名称" prop="roleName">
        <el-input v-model.trim="form.roleName" placeholder="请输入权限名称"></el-input>
      </el-form-item>
      <el-form-item label="权限描述" prop="detail">
        <el-input v-model.trim="form.detail" placeholder="请输入权限描述"></el-input>
      </el-form-item>
      <el-form-item label>
        <el-tree
          ref="tree"
          :data="auths"
          node-key="id"
          show-checkbox
          :props="defalutProps"
          :default-checked-keys="form.authIds"
          :default-expanded-keys="form.authIds"
          draggable
          @check="handleCheck"
        ></el-tree>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm('form')">立即提交</el-button>
        <el-button v-if="form.id > 0" @click="handleDialogClose">取消</el-button>
        <el-button v-else @click="resetForm('form')">重置</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script lang="ts">
import { Component, Vue, Prop, Watch } from 'vue-property-decorator'
import { editRole } from '@/api/role'
import { IAdminRole } from '@/api/types'
import { getNodes } from '@/api/auth'
import { ElForm } from 'element-ui/types/form'
import { isArray } from '@/utils/validate'

import { RoleEditDefaultForm } from './types'

@Component({
  name: 'RoleEdit'
})
export default class extends Vue {
  @Prop({
    default: false
  })
  private show!: boolean

  get dialog() {
    return this.show
  }

  set dialog(val: any) {
    this.$emit('update:show', val)
  }

  defalutProps = {
    label: 'title'
  }

  auths: any[] = []

  form = { ...RoleEditDefaultForm }

  formRule = {
    roleName: [
      { required: true, message: '请输入角色名称', trigger: 'blur' },
      { min: 2, max: 30, message: '长度在 3 到 30 个字符', trigger: 'blur' }
    ],
    detail: [
      { required: true, message: '请输入角色备注', trigger: 'blur' },
      { min: 2, max: 30, message: '长度在 3 到 30 个字符', trigger: 'blur' }
    ]
  }

  checkedAuth = ''

  reset = false

  mounted() {
    const _this = this
    ;(async function() {
      const { data } = await getNodes()
      // _this.backData = res.data.data
      // const arr = []
      // const my = res.data.data
      data.list && data.list.length && data.list.length > 0 && (_this.auths = data.list)
    })()
  }

  setForm(data: any) {
    const { authIds = [] } = data
    if (isArray(authIds)) {
      ;(this.$refs as any).tree.setCheckedKeys(authIds)
    }
    this.form = data
  }

  handleDialogClose() {
    this.reset = false
    this.$emit('update:show', false)
  }

  submitForm(formName: string) {
    ;(this.$refs[formName] as ElForm).validate(valid => {
      if (valid) {
        editRole({
          ...this.form,
          nodesData: this.checkedAuth
        }).then(() => {
          this.handleDialogClose()
          this.$emit('edit')
        })
      }
    })
  }
  resetForm(formName: string) {
    this.form = { ...RoleEditDefaultForm }
    ;(this.$refs[formName] as ElForm).clearValidate()
  }

  handleCheck(_: any, e: any) {
    const { checkedKeys = [] } = e
    checkedKeys.length && checkedKeys.length > 0 && (this.checkedAuth = checkedKeys.join(','))
  }
}
</script>

<style lang="scss" scoped>
.reset-label {
  padding-left: 10px;
  font-size: 12px;
  color: $red;
}
</style>
