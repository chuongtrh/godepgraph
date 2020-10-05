<template>
  <el-container class="container">
    <el-header>
      <el-input
        class="input-with-select"
        autocomplete="off"
        placeholder="Please input"
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
      v-loading="loading"
      element-loading-text="Loading..."
      element-loading-spinner="el-icon-loading"
      element-loading-background="rgba(0, 0, 0, 0.8)"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        xmlns:xlink="http://www.w3.org/1999/xlink"
        style="overflow: hidden"
      >
        <g class="svg-pan-zoom_viewport" v-html="svgHTML"></g>
      </svg>
    </SvgPanZoom>
  </el-container>
</template>

<script>
import axios from "axios";
import Viz from "viz.js";
import { Module, render } from "viz.js/full.render.js";
// import panzoom from "vue-panzoom";
import SvgPanZoom from "vue-svg-pan-zoom";

export default {
  components: { SvgPanZoom },
  mounted: function () {
    this.$nextTick(function () {
      this.onResize();
    });
    window.addEventListener("resize", this.onResize);
  },
  data() {
    return {
      pgkName: "net/http",
      loading: false,
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
      this.loading = true;
      let res = await axios.get(`http://localhost:3200/?pkg=${this.pgkName}`);
      let graph = res.data;
      // let graph = `digraph G { rankdir="LR"; pad=.15; ratio=auto; dpi=360; node [shape=box]; "net/http" -> "fmt"; "fmt" [style=filled,color=palegoldenrod]; "net/http" -> "io/ioutil"; "io/ioutil" [style=filled,color=palegoldenrod]; "net/http" -> "reflect"; "reflect" [style=filled,color=palegoldenrod]; "net/http" -> "sync"; "sync" [style=filled,color=palegoldenrod]; "net/http" -> "golang.org/x/net/http/httpproxy"; "golang.org/x/net/http/httpproxy" -> "net/url"; "net/url" [style=filled,color=palegoldenrod]; "golang.org/x/net/http/httpproxy" -> "os"; "os" [style=filled,color=palegoldenrod]; "golang.org/x/net/http/httpproxy" -> "strings"; "strings" [style=filled,color=palegoldenrod]; "golang.org/x/net/http/httpproxy" -> "unicode/utf8"; "unicode/utf8" [style=filled,color=palegoldenrod]; "golang.org/x/net/http/httpproxy" -> "errors"; "errors" [style=filled,color=palegoldenrod]; "golang.org/x/net/http/httpproxy" -> "fmt"; "fmt" [style=filled,color=palegoldenrod]; "golang.org/x/net/http/httpproxy" -> "net"; "net" [style=filled,color=palegoldenrod]; "net/http" -> "compress/gzip"; "compress/gzip" [style=filled,color=palegoldenrod]; "net/http" -> "os"; "os" [style=filled,color=palegoldenrod]; "net/http" -> "unicode/utf8"; "unicode/utf8" [style=filled,color=palegoldenrod]; "net/http" -> "time"; "time" [style=filled,color=palegoldenrod]; "net/http" -> "golang.org/x/net/http/httpguts"; "golang.org/x/net/http/httpguts" -> "unicode/utf8"; "unicode/utf8" [style=filled,color=palegoldenrod]; "golang.org/x/net/http/httpguts" -> "net"; "net" [style=filled,color=palegoldenrod]; "golang.org/x/net/http/httpguts" -> "net/textproto"; "net/textproto" [style=filled,color=palegoldenrod]; "golang.org/x/net/http/httpguts" -> "strings"; "strings" [style=filled,color=palegoldenrod]; "net/http" -> "net/http/httptrace"; "net/http/httptrace" [style=filled,color=palegoldenrod]; "net/http" -> "golang.org/x/net/idna"; "golang.org/x/net/idna" -> "fmt"; "fmt" [style=filled,color=palegoldenrod]; "golang.org/x/net/idna" -> "math"; "math" [style=filled,color=palegoldenrod]; "golang.org/x/net/idna" -> "strings"; "strings" [style=filled,color=palegoldenrod]; "golang.org/x/net/idna" -> "unicode/utf8"; "unicode/utf8" [style=filled,color=palegoldenrod]; "golang.org/x/net/idna" -> "golang.org/x/text/secure/bidirule"; "golang.org/x/text/secure/bidirule" -> "errors"; "errors" [style=filled,color=palegoldenrod]; "golang.org/x/text/secure/bidirule" -> "unicode/utf8"; "unicode/utf8" [style=filled,color=palegoldenrod]; "golang.org/x/net/idna" -> "golang.org/x/text/unicode/bidi"; "golang.org/x/text/unicode/bidi" -> "container/list"; "container/list" [style=filled,color=palegoldenrod]; "golang.org/x/text/unicode/bidi" -> "fmt"; "fmt" [style=filled,color=palegoldenrod]; "golang.org/x/text/unicode/bidi" -> "log"; "log" [style=filled,color=palegoldenrod]; "golang.org/x/text/unicode/bidi" -> "sort"; "sort" [style=filled,color=palegoldenrod]; "golang.org/x/text/unicode/bidi" -> "unicode/utf8"; "unicode/utf8" [style=filled,color=palegoldenrod]; "golang.org/x/net/idna" -> "golang.org/x/text/unicode/norm"; "golang.org/x/text/unicode/norm" -> "sync"; "sync" [style=filled,color=palegoldenrod]; "golang.org/x/text/unicode/norm" -> "unicode/utf8"; "unicode/utf8" [style=filled,color=palegoldenrod]; "golang.org/x/text/unicode/norm" -> "encoding/binary"; "encoding/binary" [style=filled,color=palegoldenrod]; "golang.org/x/text/unicode/norm" -> "fmt"; "fmt" [style=filled,color=palegoldenrod]; "golang.org/x/text/unicode/norm" -> "io"; "io" [style=filled,color=palegoldenrod]; "net/http" -> "crypto/tls"; "crypto/tls" [style=filled,color=palegoldenrod]; "net/http" -> "io"; "io" [style=filled,color=palegoldenrod]; "net/http" -> "log"; "log" [style=filled,color=palegoldenrod]; "net/http" -> "net"; "net" [style=filled,color=palegoldenrod]; "net/http" -> "strconv"; "strconv" [style=filled,color=palegoldenrod]; "net/http" -> "bytes"; "bytes" [style=filled,color=palegoldenrod]; "net/http" -> "container/list"; "container/list" [style=filled,color=palegoldenrod]; "net/http" -> "math/rand"; "math/rand" [style=filled,color=palegoldenrod]; "net/http" -> "mime"; "mime" [style=filled,color=palegoldenrod]; "net/http" -> "runtime"; "runtime" [style=filled,color=palegoldenrod]; "net/http" -> "crypto/rand"; "crypto/rand" [style=filled,color=palegoldenrod]; "net/http" -> "encoding/binary"; "encoding/binary" [style=filled,color=palegoldenrod]; "net/http" -> "errors"; "errors" [style=filled,color=palegoldenrod]; "net/http" -> "net/http/internal"; "net/http/internal" [style=filled,color=palegoldenrod]; "net/http" -> "net/url"; "net/url" [style=filled,color=palegoldenrod]; "net/http" -> "strings"; "strings" [style=filled,color=palegoldenrod]; "net/http" -> "sync/atomic"; "sync/atomic" [style=filled,color=palegoldenrod]; "net/http" -> "golang.org/x/net/http2/hpack"; "golang.org/x/net/http2/hpack" -> "fmt"; "fmt" [style=filled,color=palegoldenrod]; "golang.org/x/net/http2/hpack" -> "io"; "io" [style=filled,color=palegoldenrod]; "golang.org/x/net/http2/hpack" -> "sync"; "sync" [style=filled,color=palegoldenrod]; "golang.org/x/net/http2/hpack" -> "bytes"; "bytes" [style=filled,color=palegoldenrod]; "golang.org/x/net/http2/hpack" -> "errors"; "errors" [style=filled,color=palegoldenrod]; "net/http" -> "context"; "context" [style=filled,color=palegoldenrod]; "net/http" -> "math"; "math" [style=filled,color=palegoldenrod]; "net/http" -> "mime/multipart"; "mime/multipart" [style=filled,color=palegoldenrod]; "net/http" -> "net/textproto"; "net/textproto" [style=filled,color=palegoldenrod]; "net/http" -> "path"; "path" [style=filled,color=palegoldenrod]; "net/http" -> "path/filepath"; "path/filepath" [style=filled,color=palegoldenrod]; "net/http" -> "sort"; "sort" [style=filled,color=palegoldenrod]; "net/http" -> "bufio"; "bufio" [style=filled,color=palegoldenrod]; "net/http" -> "encoding/base64"; "encoding/base64" [style=filled,color=palegoldenrod]; "net/http" [style=filled,color=palegreen]; }`;

      let viz = new Viz({ Module, render });
      let svg = await viz.renderSVGElement(graph);
      let innerHTML = svg.getElementsByClassName("graph")[0].innerHTML;
      this.svgHTML = innerHTML;

      this.loading = false;
      this.center();

      // setTimeout(() => {
      //   // console.log("graph", graph);
      //   this.loading = false;
      //   this.center();
      // }, 500);
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
.el-select .el-input {
  width: 110px;
}
.input-with-select {
  background-color: #fff;
}
.container {
  position: absolute;
  overflow: auto;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  .svg-panzoom {
    background-color: #fff;
    overflow: hidden;
    // border: 1px solid black;
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
