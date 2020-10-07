<template>
  <el-footer height="20">
    <el-row :gutter="60">
      <el-col :span="16" :offset="4">
        <span>A simple tool visualize Go package - &nbsp;</span>
        <el-link
          class="github"
          type="info"
          href="https://github.com/chuongtrh/godepgraph"
          target="_blank"
        >
          Github: Godepgraph
        </el-link>
      </el-col>
    </el-row>
  </el-footer>
</template>

<style lang="scss">
.el-footer .el-col {
  display: flex;
  justify-content: center;
}
</style>
