<template>
  <el-dialog
    :title="`管理员${form.id > 0 ? '编辑' : '添加'}`"
    :visible.sync="dialog"
    @close="handleDialogClose"
    size="100%"
  >
    <el-form ref="form" :model="form" :rules="formRule" label-width="100px" style="width: 80%; margin-left:50px;">
      <el-form-item label="账户名" prop="loginName">
        <el-input v-model.trim="form.loginName" placeholder="请输入账户名"></el-input>
      </el-form-item>
      <el-form-item label="真实姓名" prop="realName">
        <el-input v-model.trim="form.realName" placeholder="请输入真实姓名"></el-input>
      </el-form-item>
      <el-form-item label="手机号" prop="phone">
        <el-input v-model.number="form.phone" placeholder="请输入手机号"></el-input>
      </el-form-item>
      <el-form-item label="邮件" prop="email">
        <el-input v-model="form.email" placeholder="请输入邮件"></el-input>
      </el-form-item>
      <el-form-item label="角色" prop="roleIds">
        <el-checkbox-group class="checkbox-group" v-model="form.roleIds">
          <el-checkbox-button style="margin-bottom: 10px;" size="medium" v-for="v in roles" :label="v.id" :key="v.id">{{
            v.title
          }}</el-checkbox-button>
        </el-checkbox-group>
      </el-form-item>
      <el-form-item v-if="form.id > 0" label="重置密码">
        <el-switch v-model="reset">重置</el-switch>
        <span class="reset-label">*密码重置123456</span>
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
import { getRoles } from '@/api/role'
import { editUser } from '@/api/users'
import { IAdminEditForm, AdminEditDefaultForm } from './types'
import { ElForm } from 'element-ui/types/form'

@Component({
  name: 'AdminEdit'
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

  form = { ...AdminEditDefaultForm }

  formRule = {
    loginName: [
      { required: true, message: '请输入登录名', trigger: 'blur' },
      { min: 3, max: 30, message: '长度在 3 到 30 个字符', trigger: 'blur' }
    ],
    realName: [
      { required: true, message: '请输入真实姓名', trigger: 'blur' },
      { min: 3, max: 30, message: '长度在 3 到 30 个字符', trigger: 'blur' }
    ],
    phone: [
      { required: true, message: '请输入手机号', trigger: 'blur' }
      // { type: 'number', message: '非法手机号', trigger: 'blur' }
    ],
    email: [
      { required: true, message: '请输入邮件', trigger: 'blur' },
      { type: 'email', message: '格式错误', trigger: 'blur' }
    ]
  }

  roles = []

  reset = false

  mounted() {
    this.getData()
  }
  async getData() {
    try {
      const res = await getRoles()
      const { data } = res
      const { list = [] } = data
      this.roles = list.map((v: any) => {
        return {
          id: v.id,
          title: v.roleName
        }
      })
    } catch (error) {}
  }

  setForm(data: any) {
    this.form = data
  }

  handleDialogClose() {
    this.reset = false
    this.$emit('update:show', false)
  }

  submitForm(formName: string) {
    ;(this.$refs[formName] as ElForm).validate(valid => {
      if (valid) {
        editUser({
          ...this.form,
          roleIds: this.form.roleIds.join(','),
          resetPwd: this.reset
        })
          .then(() => {
            this.handleDialogClose()
            this.$emit('edit')
          })
          .catch(err => {
            console.error(err)
          })
      }
    })
  }
  resetForm(formName: string) {
    this.form = { ...AdminEditDefaultForm }
    ;(this.$refs[formName] as ElForm).clearValidate()
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
