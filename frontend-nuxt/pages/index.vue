<template>
  <el-container class="container">
    <el-header class="header">
      <el-row :gutter="60">
        <el-col :span="16" :offset="4">
          <el-input
            class="search-pkg"
            autocomplete="off"
            size="large"
            clearable
            placeholder="Please input package name"
            prefix-icon="el-icon-search"
            v-model="pgkName"
            @keyup.enter="searchPgkName"
          >
            <el-button
              @click="searchPgkName"
              slot="append"
              icon="el-icon-right"
              size="mini"
            ></el-button>
          </el-input>
          <div class="example">
            Example:
            <el-link type="primary" href="/?pkg=net/http">net/http</el-link>
            <el-link type="primary" href="/?pkg=github.com/labstack/echo/v4"
              >github.com/labstack/echo/v4</el-link
            >
            <el-link type="primary" href="/?pkg=github.com/gorilla/mux"
              >github.com/gorilla/mux</el-link
            >
          </div>
        </el-col>
      </el-row>
    </el-header>
    <SvgPanZoom
      class="svg-panzoom"
      :viewportSelector="`.graph`"
      :zoomEnabled="true"
      :controlIconsEnabled="true"
      :panEnabled="true"
      :fit="false"
      :center="true"
      :minZoom="0.3"
      :maxZoom="3"
      @svgpanzoom="registerSvgPanZoom"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        xmlns:xlink="http://www.w3.org/1999/xlink"
        style="overflow: hidden"
      >
        <g class="svg-pan-zoom_viewport" v-html="svgHTML"></g>
      </svg>
    </SvgPanZoom>
    <el-link
      class="credit"
      type="info"
      href="https://github.com/chuongtrh/godepgraph"
    >
      Repo Github: Godepgraph
    </el-link>
  </el-container>
</template>

<script>
import Viz from "viz.js";
import { Module, render } from "viz.js/full.render.js";
import SvgPanZoom from "vue-svg-pan-zoom";
export default {
  components: { SvgPanZoom },
  mounted: function () {
    this.$nextTick(function () {
      this.onResize();
    });
    window.addEventListener("resize", this.onResize);

    this.pgkName = this.$route.query.pkg || "";
    console.log("pgkName", this.pgkName);
    if (this.pgkName != "") {
      this.searchPgkName();
    }
  },
  data() {
    return {
      pgkName: "net/http",
      svgHTML: `<g></g>`,
      svgpanzoom: null,
    };
  },
  methods: {
    onResize() {
      this.resize();
    },
    async searchPgkName() {
      console.log("Search", this.pgkName);

      const loading = this.$loading({
        lock: true,
        text: "Fetch package ...",
        spinner: "el-icon-loading",
        background: "rgba(0, 0, 0, 0.7)",
      });

      try {
        let res = await this.$axios.get(
          `${process.env.baseUrl}/search/?pkg=${this.pgkName}`
        );
        if (res.status == 200) {
          let graph = res.data;
          let viz = new Viz({ Module, render });
          loading.text = "Loading graph...";
          let svg = await viz.renderSVGElement(graph);
          let innerHTML = svg.getElementsByClassName("graph")[0].innerHTML;
          this.svgHTML = innerHTML;

          this.center();
          setTimeout(() => {
            loading.close();
          }, 300);
        } else {
          loading.close();
          console.log(res.data);
        }
      } catch (error) {
        loading.close();
        this.svgHTML = `<g></g>`;
        this.$message.error(`Oops, package ${this.pgkName} not found!`);
      }
    },
    registerSvgPanZoom(svgpanzoom) {
      this.svgpanzoom = svgpanzoom;
    },
    resize() {
      if (!this.svgpanzoom) return;

      this.svgpanzoom.resize();
    },
    center() {
      if (!this.svgpanzoom) return;

      this.svgpanzoom.center();
    },
  },
};
</script>

<style lang="scss">
.container {
  position: absolute;
  overflow: hidden;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  .header {
    height: 120px;
    margin-top: 1rem;
    .search-pkg {
      background-color: #fff;

      .el-select .el-input {
        width: 110px;
      }
    }
    .example {
      margin-top: 0.5rem;
      margin-bottom: 0.5rem;
      display: flex;
      font-size: 12px;
      font-weight: 400;
      a {
        font-size: 12px;
        font-weight: 400;
        margin-left: 1rem;
      }
    }
  }

  .svg-panzoom {
    background-color: #fff;
    overflow: hidden;
    margin-top: 1rem;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    svg {
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
      overflow: hidden;
    }
  }
}
</style>
