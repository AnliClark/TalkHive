<template>
  <div 
    class="editable-text" 
    @mouseover="showEditIcon = true" 
    @mouseleave="showEditIcon = false"
    @click="startEditing"
  >
    <div v-if="!isEditing" class="text-display">
      <span>{{ text }}</span>
      <i v-if="showEditIcon" class="edit-icon" >✏️</i>
    </div>
    <div v-else class="text-edit">
      <input 
        type="text" 
        v-model="editableText" 
        @blur="saveEdit" 
        @keyup.enter="saveEdit" 
        ref="editInput"
      />
    </div>
  </div>
</template>

<script>
export default {
  props: {
    text: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      isEditing: false,
      showEditIcon: false,
      editableText: this.text
    };
  },
  methods: {
    startEditing() {
      this.isEditing = true;
      this.$nextTick(() => {
        this.$refs.editInput.focus();
      });
    },
    saveEdit() {
      if(!this.isEditing) return;  // 防止被重复调用
      this.isEditing = false;
      this.$emit('update-text', this.editableText);
    }
  },
  watch: {
    text(newText) {
      this.editableText = newText;
    }
  },
};
</script>

<style scoped>
.editable-text {
  position: relative;
  cursor: pointer;
}

.text-display {
  display: flex;
  align-items: flex-start;
}

.edit-icon {
  margin-left: 5px;
  cursor: pointer;
  font-size: var(--font-size-small);
}

.text-edit input {
  width: 100%;
  padding: 5px;
  font-size: var(--font-size-small);
}

span {
  text-align: left;
}
</style>