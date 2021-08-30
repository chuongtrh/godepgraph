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
            @change="searchPgkName"
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
            <el-link type="primary" href="/?pkg=github.com/gofiber/fiber/v2"
              >github.com/gofiber/fiber/v2</el-link
            >
            <el-link type="primary" href="/?pkg=github.com/labstack/echo/v4"
              >github.com/labstack/echo/v4</el-link
            >
            <el-link type="primary" href="/?pkg=go.uber.org/zap"
              >go.uber.org/zap</el-link
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
      :minZoom="0.1"
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
    <Footer />
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
    if (this.pgkName != "") {
      this.searchPgkName();
    }
  },
  data() {
    return {
      pgkName: "net/http",
      svgHTML: `<g></g>`,
      svgpanzoom: null,
      viewBox: null,
    };
  },
  methods: {
    onResize() {
      this.resize();
    },
    async searchPgkName() {
      // TODO: update URL
      this.$router.push({
        path: this.$route.path,
        query: { pkg: this.pgkName },
      });
      if (!this.pgkName) {
        return;
      }
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
          let svg = await viz.renderSVGElement(graph, {
            totalMemory: 32 * 1024 * 1024,
          });
          this.viewBox = svg.getAttribute("viewBox").split(" ").map(Math.ceil);
          let innerHTML = svg.getElementsByClassName("graph")[0].innerHTML;
          this.svgHTML = innerHTML;
          this.fitCenter();

          setTimeout(() => {
            loading.close();
          }, 100);
        } else {
          loading.close();
          console.error(res.data);
        }
      } catch (error) {
        loading.close();
        this.svgHTML = `<g></g>`;
        this.viewBox = null;
        this.$message({
          message: `Oops, package ${this.pgkName} not found!`,
          type: "error",
          offset: 120,
        });
      }
    },
    registerSvgPanZoom(svgpanzoom) {
      this.svgpanzoom = svgpanzoom;
    },
    resize() {
      if (!this.svgpanzoom) return;

      this.svgpanzoom.resize();
    },
    fitCenter() {
      if (!this.svgpanzoom || !this.viewBox) return;

      let width = this.viewBox[2];
      let height = this.viewBox[3];
      var elmnt = window.document.querySelector(".svg-panzoom");
      if (elmnt) {
        let clientWidth = elmnt.clientWidth;
        let clientHeight = elmnt.clientHeight;
        this.svgpanzoom.pan({
          x: Math.abs(clientWidth - width) / 2,
          y: Math.abs(height),
        });
      }
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
      flex-wrap: wrap;
      font-size: 12px;
      font-weight: 400;
      a {
        font-size: 12px;
        font-weight: 400;
        margin-left: 1rem;
      }
    }
    .el-input-group__append {
      background-color: #409eff;
      .el-icon-right {
        color: #fff;
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
