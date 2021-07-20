<template>
  <multipane class="custom-resizer" layout="vertical">
    <div class="pane menu">
      <div class="avatar">
        <el-avatar
            shape="square"
            size="large"
            :src="userInfo.avatar_url"
        ></el-avatar>
        <h2>{{ userInfo.name }}</h2>
        <el-button type="primary" @click="createNote">
          新建笔记
        </el-button>
      </div>
    </div>
    <multipane-resizer></multipane-resizer>
    <div class="pane doc" :style="{ flexGrow: 1 }">
      <div v-show="activeTreeNode">
        <div class="doc-title">
          <el-input v-model="activeTreeNode.name"></el-input>
          <el-button @click="updateNote">保存</el-button>
        </div>
        <div id="vditor"></div>
      </div>
      <el-empty v-show="!activeTreeNode"></el-empty>
    </div>
  </multipane>
</template>

<script lang="ts">
import {ref, defineComponent, onMounted, reactive, toRefs} from "vue";
import {useRouter, useRoute} from "vue-router";
import service from "@/utils/request";
import {Multipane, MultipaneResizer} from "@/components/Multipane";
import {getNowDate} from "@/utils/date";
import Vditor from "vditor";

export default defineComponent({
  name: "App",
  components: {
    Multipane,
    MultipaneResizer,
  },
  data() {
    return {
      userInfo: {
        name: '',
        avatar_url: ''
      },
      activeTreeNode: {
        name: ''
      }
    }
  },
  methods: {
    createNote() {

    },
    updateNote() {

    }
  }
});
</script>

<style scoped lang="scss">
.custom-resizer {
  width: 100%;
  height: 100vh;
}

.custom-resizer > .pane {
  text-align: left;
  padding: 15px;
  overflow: hidden;
  background: #fff;
  border: 1px solid #ccc;

  &.menu {
    width: 20%;
    max-width: 250px;
  }

  &.catalogue {
    width: 300px;
    background-color: #f4f6f9;
  }
}

.custom-resizer > .multipane-resizer {
  margin: 0;
  left: 0;
  position: relative;

  &:before {
    display: block;
    content: "";
    width: 3px;
    height: 40px;
    position: absolute;
    top: 50%;
    left: 50%;
    margin-top: -20px;
    margin-left: -1.5px;
    border-left: 1px solid #ccc;
    border-right: 1px solid #ccc;
  }

  &:hover {
    &:before {
      border-color: #999;
    }
  }
}

.menu {
  .avatar {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    margin-bottom: 20px;
  }
}

.catalogue {
  .ant-divider-horizontal {
    margin: 15px 0;
  }

  .info {
    display: flex;
    align-items: center;
    justify-content: space-between;

    p {
      margin: 0;
    }
  }
}

.doc {
  &-title {
    display: flex;
    margin-bottom: 10px;
  }
}
</style>
