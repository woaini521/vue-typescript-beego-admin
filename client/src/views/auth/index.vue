<template>
  <div class="dashboard-container">
    <el-row>
      <el-col :sm="24" :md="8" style="overflow-x: scrollbar;">
        <el-tree
          :data="auths"
          node-key="id"
          :props="defalutProps"
          :render-content="renderContent"
          :expand-on-click-node="false"
          draggable
          @node-drop="treeSort"
          @node-click="treeClick"
          :allow-drop="treeCanSort"
        ></el-tree>
      </el-col>
      <el-col :sm="24" :md="16">
        <el-form
          :model="formAuth"
          :rules="formAuth.isShow ? rulesWeb : rules"
          ref="authForm"
          label-width="80px"
          style="width: 500px"
        >
          <el-form-item label="权限名称" prop="title">
            <el-input v-model="formAuth.title" type="text" clearable></el-input>
          </el-form-item>
          <el-form-item label="后端路由" prop="authUrl">
            <el-input v-model="formAuth.authUrl" type="text"></el-input>
          </el-form-item>
          <el-form-item label="菜单图标" prop="icon">
            <el-input v-model="formAuth.icon" type="text"></el-input>
          </el-form-item>
          <el-form-item label="前端配置" prop="isShow">
            <el-switch
              v-model="formAuth.isShow"
              active-color="#13ce66"
              inactive-color="#ff4949"
              :active-value="1"
              :inactive-value="0"
            ></el-switch>
          </el-form-item>
          <el-form-item v-if="formAuth.isShow" label="前端路由" prop="path">
            <el-input v-model="formAuth.path" type="text"></el-input>
          </el-form-item>
          <el-form-item v-if="formAuth.isShow" label="重定向" prop="redirect">
            <el-input v-model="formAuth.redirect" type="text"></el-input>
          </el-form-item>
          <el-form-item v-if="formAuth.isShow" label="前端组件" prop="component">
            <el-select v-model="formAuth.component" placeholder="一级菜单路由不选择" filterable>
              <el-option
                v-for="component in selectList"
                :key="component.key"
                :value="component.key"
                :label="component.title"
              ></el-option>
            </el-select>
            <!-- <el-input v-model="formAuth.component" type="text"></el-input> -->
          </el-form-item>
          <el-form-item v-if="formAuth.isShow" label="隐藏菜单" prop="slideebarHidden">
            <el-switch
              v-model="formAuth.slideebarHidden"
              active-color="#13ce66"
              inactive-color="#ff4949"
              :active-value="1"
              :inactive-value="0"
            ></el-switch>
          </el-form-item>
          <el-form-item v-if="formAuth.isShow" label="面包屑" prop="breadcrumb">
            <el-switch
              v-model="formAuth.breadcrumb"
              active-color="#13ce66"
              inactive-color="#ff4949"
              :active-value="1"
              :inactive-value="0"
            ></el-switch>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitForm('authForm')">提交</el-button>
            <el-button @click="resetForm('authForm')">重置</el-button>
          </el-form-item>
        </el-form>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import { UserModule } from '@/store/modules/user'
import { CreateElement } from 'vue'
import { getNodes, editNodes } from '@/api/auth'
import { authCompontsConfig, authComponents } from '@/router/authRouter'
import { ElForm } from 'element-ui/types/form'

const defalutFormAuth = {
  id: -1,
  title: '',
  authUrl: '',
  pid: 0,
  sort: 0,
  icon: '',
  isShow: '',
  path: '',
  slideebarHidden: 0,
  breadcrumb: 1,
  component: '',
  redirect: ''
}

const validateComponent = (rule: any, value: any, cb: Function) => {
  if (!value) {
    return cb()
  }
  if (value && !Object.prototype.hasOwnProperty.call(authComponents, value)) {
    return cb(new Error('组件不存在'))
  }
  cb()
}

@Component({
  name: 'Auth'
})
export default class extends Vue {
  defalutProps = {
    label: 'title'
  }
  auths: any[] = []
  formAuth = {
    ...defalutFormAuth
  }

  selectList = authCompontsConfig()

  rules = {
    title: [
      { required: true, message: '请输入权限名称', trigger: 'blur' },
      { min: 2, max: 10, message: '长度在 2 到 10 个字符', trigger: 'blur' }
    ]
  }
  rulesWeb = {
    ...this.rules,
    path: [{ required: true, message: '请输入页面路由', trigger: 'blur' }],
    component: [{ validator: validateComponent, trigger: 'blur' }]
  }

  //
  treeClick(data: any, node: any, env: any) {
    data.id !== 1 && (this.formAuth = data)
  }
  treeCanSort(draggingNode: any, dropNode: any, type: any) {
    if (draggingNode.data.id === 0) {
      return false
    }
    if (draggingNode.data.pId !== dropNode.data.pId && type !== 'inner') {
      return false
    }
    return true
  }
  treeSort(draggingNode: any, dropNode: any, dropType: any, ev: any) {
    let data = []
    if (dropType === 'inner') {
      data = dropNode.childNodes.map((v: any, i: number) => {
        v.data['sort'] = i
        v.data['pId'] = dropNode.data.id
        return v.data
      })
    } else {
      data = dropNode.parent.childNodes.map((v: any, i: number) => {
        v.data['sort'] = i
        v.data['pId'] = dropNode.parent.data.id
        return v.data
      })
    }
    editNodes({
      list: JSON.stringify(
        data.map((v: any) => {
          v.children = null
          return v
        })
      )
    })
  }
  treeNodeEdit(data: any) {
    console.log(data)
    // console.log(data)
    // this.formAuth = data
    window.event && window.event.stopPropagation()
  }
  treeNodeAppend(data: any) {
    this.formAuth = {
      ...defalutFormAuth,
      id: -1,
      pid: data.pid,
      title: '上级权限' + data.title
    }
    window.event && window.event.stopPropagation()

    // const newChild = { id: 500,  title:'add', children: [] }
    // if (!data.children) {
    //   this.$set(data, 'children', [])
    // }
    // data.children.push(newChild)
  }
  treeNodeRemove(node: any, data: any) {
    const parent = node.parent
    const children = parent.data.children || parent.data
    const index = children.findIndex((d: any) => d.id === data.id)
    children.splice(index, 1)
    window.event && window.event.stopPropagation()
  }
  renderContent(h: CreateElement, { node, data, store }: any) {
    return h('span', [
      h('span', node.label),
      h('span', [
        h(
          'el-button',
          {
            style: { paddingLeft: '20px' },
            attrs: { size: 'mini', type: 'text', 'data-node': data },
            on: { click: this.treeNodeEdit.bind(this, data) }
          },
          '编辑'
        ),
        h(
          'el-button',
          {
            attrs: { size: 'mini', type: 'text' },
            on: { click: this.treeNodeRemove.bind(this, node, data) }
          },
          '删除'
        ),
        h(
          'el-button',
          {
            attrs: { size: 'mini', type: 'text' },
            on: { click: this.treeNodeAppend.bind(this, data) }
          },
          '增加'
        )
      ])
    ])
    // return (
    //   <span class="custom-tree-node">
    //     <span>{node.label}</span>
    //     <span>
    //       <el-button size="mini" type="text" on-click={() => this.append(data)}>
    //         Append
    //       </el-button>
    //       <el-button
    //         size="mini"
    //         type="text"
    //         on-click={() => this.remove(node, data)}
    //       >
    //         Delete
    //       </el-button>
    //     </span>
    //   </span>
    // )
  }

  submitForm(formName: string) {
    ;(this.$refs[formName] as ElForm).validate(valid => {
      if (valid) {
        // console.error({
        //   list: JSON.stringify([{ ...this.formAuth }])
        // })
        editNodes({
          list: JSON.stringify([{ ...this.formAuth }])
        })
      } else {
        console.log('error submit!!')
        return false
      }
    })
  }
  resetForm(formName: string) {
    this.formAuth = { ...defalutFormAuth }
    ;(this.$refs[formName] as ElForm).clearValidate()
  }

  mounted() {
    try {
      const _this = this
      ;(async function() {
        const { data } = await getNodes()
        // _this.backData = res.data.data
        // const arr = []
        // const my = res.data.data
        data.list && data.list.length && data.list.length > 0 && (_this.auths = data.list)
      })()
    } catch (error) {}
  }
}
</script>

<style lang="scss" scoped>
.dashboard {
  &-container {
    margin: 30px;
  }

  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>
