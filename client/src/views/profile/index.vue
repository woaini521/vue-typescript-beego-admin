<template>
  <div class="app-container">
    <el-form ref="form" :model="form" :rules="rules" label-width="100px" style="width: 500px;">
      <el-form-item label="登录名：" prop="loginName">
        <el-input disabled v-model="form.loginName" />
      </el-form-item>
      <el-form-item label="真实姓名：" prop="realName">
        <el-input :disabled="form.id === 1" v-model="form.realName" />
      </el-form-item>
      <el-form-item label="手机号：" prop="phone">
        <el-input v-model="form.phone" />
      </el-form-item>
      <el-form-item label="邮件：" prop="email">
        <el-input v-model="form.email" />
      </el-form-item>
      <el-form-item label="修改密码">
        <el-switch v-model="form.resetPwd" :active-value="true"></el-switch>
      </el-form-item>
      <el-form-item v-if="form.resetPwd" label="密码" prop="passwordOld">
        <el-input type="password" v-model="form.passwordOld" placeholder="请输入密码"></el-input>
      </el-form-item>
      <el-form-item v-if="form.resetPwd" label="新密码" prop="password">
        <el-input type="password" v-model="form.password" placeholder="请输入新密码"></el-input>
      </el-form-item>
      <el-form-item v-if="form.resetPwd" label="确认密码" prop="passwordSure">
        <el-input type="password" v-model="form.passwordSure" placeholder="请输入新密码"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm('form')">
          立即提交
        </el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { getUserInfo, editUser } from '@/api/users'

import { ElForm } from 'element-ui/types/form'

@Component({
  name: 'ProfileEdit'
})
export default class extends Vue {
  private form = {
    id: -1,
    loginName: '',
    realName: '',
    phone: '',
    email: '',
    resetPwd: false,
    passwordOld: '',
    password: '',
    passwordSure: ''
  }
  private validatePasswordSure(rule: any, value: string, callback: Function) {
    if (!value) {
      callback(new Error('请输入确认密码'))
    } else if (value !== this.form.password) {
      callback(new Error('密码不一致'))
    } else {
      callback()
    }
  }
  private rules = {
    loginName: [{ required: true, message: '请输入登录名', trigger: 'blur' }],
    realName: [{ required: true, message: '请输入真实姓名名', trigger: 'blur' }],
    phone: [{ required: true, message: '请输入手机号', trigger: 'blur' }],
    email: [
      { required: true, message: '请输入邮件', trigger: 'blur' },
      { type: 'email', message: '格式错误', trigger: 'blur' }
    ],
    passwordOld: [{ required: true, message: '请输入密码', trigger: 'blur' }],
    password: [{ required: true, message: '请输入新密码', trigger: 'blur' }],
    passwordSure: [
      { required: true, message: '请输入邮件', trigger: 'blur' },
      { validator: this.validatePasswordSure, trigger: 'blur' }
    ]
  }

  mounted() {
    this.getData()
  }

  getData() {
    getUserInfo({ profile: 1 }).then(res => {
      const { data } = res
      Object.prototype.toString.call(data) === '[object Object]' && (this.form = { ...this.form, ...data })
    })
  }

  private submitForm(formName: string) {
    ;(this.$refs[formName] as ElForm).validate(valid => {
      if (valid) {
        editUser({
          ...this.form,
          profile: true
        }).then(() => this.getData())
      }
    })
  }

  private onCancel() {
    this.$message({
      message: 'cancel!',
      type: 'warning'
    })
  }
}
</script>

<style lang="scss" scoped>
.line {
  text-align: center;
}
</style>
