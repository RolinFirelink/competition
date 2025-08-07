<template>
  <div class="section">
    <h2 class="section-title">对局投稿</h2>
    
    <!-- 投稿表单 -->
    <form @submit.prevent="submitForm" class="submission-form">
      <div class="form-group">
        <label class="form-label">对局类型 *</label>
        <select 
          v-model="formData.type" 
          class="form-select" 
          required
        >
          <option value="">请选择对局类型</option>
          <option value="超唐对局">超唐对局</option>
          <option value="精彩对局">精彩对局</option>
          <option value="上大师了!">上大师了!</option>
        </select>
      </div>

      <div class="form-group">
        <label class="form-label">录像ID *</label>
        <input 
          type="text" 
          v-model="formData.replayId" 
          class="form-input" 
          placeholder="请输入录像ID"
          required
        />
      </div>

      <div class="form-group">
        <label class="form-label">对局描述</label>
        <textarea 
          v-model="formData.description" 
          class="form-textarea" 
          :placeholder="descriptionPlaceholder"
        ></textarea>
      </div>

      <div class="form-actions">
        <button type="button" @click="resetForm" class="btn btn-secondary">
          撤销
        </button>
        <button type="submit" class="btn">
          提交
        </button>
      </div>
    </form>

    <!-- 帮助信息 -->
    <div class="help-section">
      <h4>如何找到对局ID？</h4>
      <p>在街霸6游戏中，对局结束后可以在回放系统中找到对应的录像ID。</p>
      <a href="#" class="help-link">查看详细教程 →</a>
    </div>

    <!-- 成功提示 -->
    <div v-if="showSuccess" class="success-message">
      投稿提交成功！感谢您的分享。
    </div>
  </div>
</template>

<script>
export default {
  name: 'SubmissionForm',
  data() {
    return {
      formData: {
        type: '',
        replayId: '',
        description: ''
      },
      showSuccess: false
    }
  },
  computed: {
    descriptionPlaceholder() {
      return this.formData.type === '上大师了!' 
        ? '请描述重点内容在第几局,方便UP主迅速定位内容,合适的描述将让您有更高概率被选入~ 请在描述中注明您的抖音或B站ID'
        : '请描述重点内容在第几局,方便UP主迅速定位内容,合适的描述将让您有更高概率被选入~'
    }
  },
  methods: {
    submitForm() {
      // 验证表单
      if (!this.formData.type || !this.formData.replayId) {
        alert('请填写必填项')
        return
      }

      // 验证粉丝福利投稿格式
      if (this.formData.type === '上大师了!' && this.formData.description) {
        const hasSocialMediaId = /(抖音|b站|B站|douyin|bilibili)/i.test(this.formData.description)
        if (!hasSocialMediaId) {
          alert('粉丝福利投稿需要在描述中注明您的抖音或B站ID')
          return
        }
      }

      // 创建投稿数据
      const submission = {
        id: Date.now(),
        type: this.formData.type,
        replayId: this.formData.replayId,
        description: this.formData.description || '无描述',
        timestamp: new Date().toISOString(),
        status: 'pending'
      }

      // 发送到全局事件总线（模拟提交到后台）
      this.$emit('submission-submitted', submission)

      // 显示成功提示
      this.showSuccess = true
      setTimeout(() => {
        this.showSuccess = false
      }, 3000)

      // 重置表单
      this.resetForm()
    },

    resetForm() {
      this.formData = {
        type: '',
        replayId: '',
        description: ''
      }
    }
  }
}
</script>

<style scoped>
.submission-form {
  margin-bottom: 30px;
}

.form-actions {
  display: flex;
  gap: 15px;
  justify-content: flex-end;
  margin-top: 20px;
}

.help-section {
  background-color: var(--light-background);
  padding: 20px;
  border-radius: var(--border-radius);
  border-left: 4px solid var(--accent-color);
}

.help-section h4 {
  color: var(--primary-color);
  margin-bottom: 10px;
  font-size: 16px;
}

.help-section p {
  color: var(--text-color);
  margin-bottom: 10px;
  font-size: 14px;
}

.help-link {
  color: var(--accent-color);
  text-decoration: none;
  font-size: 14px;
  font-weight: 500;
}

.help-link:hover {
  text-decoration: underline;
}

@media (max-width: 768px) {
  .form-actions {
    flex-direction: column;
  }
  
  .form-actions .btn {
    width: 100%;
  }
}
</style> 