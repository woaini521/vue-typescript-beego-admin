<template>
  <div :class="{hidden: hidden}" class="pagination-container">
    <el-pagination
      :background="background"
      :current-page.sync="currentPage"
      :page-size.sync="pageSize"
      :layout="layout"
      :page-sizes="pageSizes"
      :total="total"
      v-bind="$attrs"
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
    />
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'

@Component({
  name: 'Pagination'
})
export default class extends Vue {
  @Prop({ required: true }) private total!: Number
  @Prop({ default: 1 }) private page!: Number
  @Prop({ default: 20 }) private limit!: Number
  @Prop({ default: () => [1, 2, 10, 20, 30, 50, 70, 100] }) private pageSizes!: Number[]
  @Prop({ default: 'total, sizes, prev, pager, next, jumper' }) private layout!: String
  @Prop({ default: true }) private background!: Boolean
  @Prop({ default: true }) private autoScroll!: Boolean
  @Prop({ default: false }) private hidden!: Boolean

  get currentPage() {
    return this.page
  }

  get pageSize() {
    return this.limit
  }

  set currentPage(val: any) {
    this.$emit('update:page', val)
  }

  set pageSize(val: any) {
    this.$emit('update:page', 1)
    this.$emit('update:limit', val)
  }

  private handleSizeChange(val: number) {
    this.$emit('pagination', { page: 1, limit: val })
    if (this.autoScroll) {
      scrollTo(0, 800)
    }
  }
  private handleCurrentChange(val: number) {
    this.$emit('pagination', { page: val, limit: this.pageSize })
    if (this.autoScroll) {
      scrollTo(0, 800)
    }
  }
}
</script>

<style lang="scss" scoped>
.pagination-container {
  background: #fff;
  padding: 32px 16px;
}
.pagination-container.hidden {
  display: none;
}
</style>
