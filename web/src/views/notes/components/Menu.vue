<template>
  <div class="avatar">
    <a-avatar shape="square" size="large" :src="userInfo.avatar_url"></a-avatar>
    <h2>{{ userInfo.name }}</h2>
    <a-button type="primary" @click="createNote">
      新建笔记
      <PlusOutlined/>
    </a-button>
  </div>
  <div>
    <p>近期笔记</p>
    <p>我的分享</p>
    <a-directory-tree :tree-data="workspace.note"
                      v-model:selectedKeys="selectedKeys"
                      :replace-fields="replaceFields"
                      @select="selectTreeNode"
    >
      <template #title="node">
        <a-dropdown :trigger="['contextmenu']">
          <a-input v-if="node.isEdit" :value="node.name" @keyup.enter="updateName(node)"></a-input>
          <span v-else>
                <span v-if="node.isDir">{{ node.name }}</span>
              </span>
          <template #overlay>
            <a-menu @click="({ key: menuKey }) => onContextMenuClick(node, menuKey)">
              <a-menu-item key="1">新建子文件夹</a-menu-item>
              <a-menu-item key="2">重命名</a-menu-item>
              <a-menu-item key="3">删除</a-menu-item>
              <a-menu-item key="4">向上移动</a-menu-item>
              <a-menu-item key="5">向下移动</a-menu-item>
            </a-menu>
          </template>
        </a-dropdown>
      </template>
    </a-directory-tree>
  </div>
</template>

<script lang="ts">
import {defineComponent, onMounted, ref} from 'vue'
import {SearchOutlined, UnorderedListOutlined, PlusOutlined} from '@ant-design/icons-vue';

export default defineComponent({
  name: "Menu",
  components: {SearchOutlined, UnorderedListOutlined, PlusOutlined},
  setup() {
    const data = receive({
      userInfo: {
        avatar_url: '',
        name: ''
      },
      workspace: {
        note: []
      }
    })
    const selectedKeys = ref<string[]>([]);
    const replaceFields = {
      children: 'child',
      title: 'name',
      key: 'path'
    }

    //  新建笔记
    function createNote() {
      const dataRef = data.activeNotes[0].dataRef
      data.notes.push({
        "name": "无标题",
        "path": `${dataRef.path}/无标题.md`,
        "isDir": false,
        "isShare": false,
        "sha": "",
        "size": 0,
        "url": "",
        "type": "",
        "html_url": "",
        "download_url": "",
        "createTime": getNowDate(),
        "updateTime": getNowDate(),
      })
      data.activeNotes[0].dataRef.child = data.notes;
    }

    return {
      selectedKeys,
      replaceFields
    }
  }
})
</script>

<style scoped>

</style>