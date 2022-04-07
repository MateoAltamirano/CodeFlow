<template>
  <div class="flow">
    <el-aside width="250px" class="column">
      <el-button color="#00cd6a" size="large" @click="runProgram"
        >Run</el-button
      >
      <el-button color="#0099de" size="large" @click="saveDialogVisible = true"
        >Save</el-button
      >
      <el-button color="#ffc500" size="large" @click="openImportDialog"
        >Import</el-button
      >
    </el-aside>
    <el-main class="code">
      <div class="code-generated">
        <pre>
          <p style="color: white">{{ codeFlowStore.code }}</p>
        </pre>
      </div>
      <div class="code-executed">
        <pre>
          <p style="color: white">{{ codeFlowStore.result }}</p>
        </pre>
      </div>
    </el-main>
    <el-dialog
      v-model="importDialogVisible"
      title="Choose the program"
      width="70%"
    >
      <el-select v-model="programId" placeholder="Program">
        <el-option
          v-for="item in options"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        />
      </el-select>
      <template #footer>
        <span class="dialog-footer">
          <el-button
            color="#ff4b3a"
            style="height: 3rem; margin-right: 1rem"
            @click="importDialogVisible = false"
            >Cancel</el-button
          >
          <el-button color="#0099de" style="height: 3rem" @click="importProgram"
            >Import</el-button
          >
        </span>
      </template>
    </el-dialog>
    <el-dialog v-model="saveDialogVisible" title="Save the program" width="70%">
      <el-input v-model="programName" />
      <template #footer>
        <span class="dialog-footer">
          <el-button
            color="#ff4b3a"
            style="height: 3rem; margin-right: 1rem"
            @click="saveDialogVisible = false"
            >Cancel</el-button
          >
          <el-button color="#0099de" style="height: 3rem" @click="saveProgram"
            >Save</el-button
          >
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import { useCodeFlowStore } from '../stores/codeFlow';
import { inject, ref } from 'vue';

export default {
  setup() {
    const importDialogVisible = ref(false);
    const saveDialogVisible = ref(false);
    const codeFlowStore = useCodeFlowStore();
    const axios = inject('axios');
    const programId = ref(null);
    const programName = ref('');
    const options = [];

    const runProgram = () => {
      axios
        .post('http://localhost:9000/programs/execute', {
          Code: codeFlowStore.code,
        })
        .then((response) => {
          codeFlowStore.updateStore('result', response.data);
        })
        .catch((error) => {
          console.log(error);
        });
    };

    const saveProgram = () => {
      if (programName.value.trim() === '') {
        alert('Please enter a name');
        return;
      }
      axios
        .post('http://localhost:9000/programs', {
          Name: programName.value.trim(),
          Code: codeFlowStore.code,
          Flow: JSON.stringify(codeFlowStore.flow),
        })
        .then((response) => {
          programName.value = '';
        })
        .catch((error) => {
          console.log(error);
        });
      saveDialogVisible.value = false;
    };

    const openImportDialog = () => {
      axios
        .get(`http://localhost:9000/programs`)
        .then((response) => {
          const { programs } = response.data;
          programs.forEach((program) => {
            options.push({
              value: program.uid,
              label: program.name,
            });
          });
          importDialogVisible.value = true;
        })
        .catch((error) => {
          console.log(error);
        });
    };

    const importProgram = () => {
      if (!programId.value) {
        alert('Please choose a program');
        return;
      }
      axios
        .get(`http://localhost:9000/programs/${programId.value}`)
        .then((response) => {
          const { program } = response.data;
          codeFlowStore.updateStore('code', program[0].code);
          codeFlowStore.updateStore('flow', JSON.parse(program[0].flow));
        })
        .catch((error) => {
          console.log(error);
        });
      importDialogVisible.value = false;
    };

    return {
      codeFlowStore,
      runProgram,
      saveProgram,
      importProgram,
      importDialogVisible,
      saveDialogVisible,
      options,
      programId,
      programName,
      openImportDialog,
    };
  },
};
</script>

<style>
.flow {
  display: flex;
  height: 100%;
}

.column {
  display: flex;
  flex-direction: column;
  overflow: auto;
  background-color: #23475f;
  padding: 1rem;
}

.code {
  display: flex;
  height: 100%;
  flex-direction: column;
  background-color: #1e3c50;
}

.code-generated {
  border: solid 1px #798d8e;
  border-radius: 4px;
  display: flex;
  height: 100%;
  margin-bottom: 0.5rem;
  overflow: auto;
}

.code-executed {
  border: solid 1px #798d8e;
  border-radius: 4px;
  display: flex;
  height: 100%;
  margin-top: 0.5rem;
  overflow: auto;
}

pre {
  padding: 0 1rem;
  tab-size: 4;
}

.el-button {
  height: 60px;
  line-height: 40px;
  border-radius: 0.5rem;
  margin-bottom: 1rem;
  color: white;
}
.el-button:hover {
  color: white;
}
.el-button:focus {
  color: white;
}
.el-button + .el-button {
  margin-left: 0;
}
</style>
