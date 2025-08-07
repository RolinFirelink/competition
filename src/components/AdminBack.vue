<template>
  <div class="container">
    <div class="admin-header">
      <h1 class="admin-title">每日一灯后台管理</h1>
      <p class="admin-subtitle">投稿对局数据管理</p>
    </div>

    <div class="admin-content">
      <div class="stats-section">
        <div class="stat-card">
          <div class="stat-number">{{ submissions.length }}</div>
          <div class="stat-label">总投稿数</div>
        </div>
        <div class="stat-card">
          <div class="stat-number">{{ pendingCount }}</div>
          <div class="stat-label">待处理</div>
        </div>
        <div class="stat-card">
          <div class="stat-number">{{ todayCount }}</div>
          <div class="stat-label">今日投稿</div>
        </div>
      </div>

      <div class="submissions-section">
        <h2 class="section-title">投稿列表</h2>
        
        <div v-if="submissions.length === 0" class="empty-state">
          <div class="empty-icon">📝</div>
          <p>暂无投稿数据</p>
          <p class="empty-hint">当前没有收到任何对局投稿</p>
        </div>

        <div v-else class="submissions-list">
          <div 
            v-for="submission in sortedSubmissions" 
            :key="submission.id" 
            class="submission-item"
          >
            <div class="submission-header">
              <div class="submission-type" :class="getTypeClass(submission.type)">
                {{ submission.type }}
              </div>
              <div class="submission-time">
                {{ formatTime(submission.timestamp) }}
              </div>
            </div>
            
            <div class="submission-content">
              <div class="submission-field">
                <label>录像ID:</label>
                <span class="field-value">{{ submission.replayId }}</span>
              </div>
              
              <div class="submission-field">
                <label>对局描述:</label>
                <span class="field-value">{{ submission.description }}</span>
              </div>
              
              <div class="submission-field">
                <label>状态:</label>
                <span class="status-badge" :class="getStatusClass(submission.status)">
                  {{ getStatusText(submission.status) }}
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'AdminBack',
  data() {
    return {
      submissions: [
        // 示例数据，实际应该从全局状态或localStorage获取
        {
          id: 1,
          type: '超唐对局',
          replayId: 'SF6_20241201_001',
          description: '第三局关键时刻的连招，非常精彩',
          timestamp: '2024-12-01T10:30:00.000Z',
          status: 'pending'
        },
        {
          id: 2,
          type: '上大师了!',
          replayId: 'SF6_20241201_002',
          description: '终于上大师了！我的抖音ID是：fighting_master',
          timestamp: '2024-12-01T09:15:00.000Z',
          status: 'pending'
        },
        {
          id: 3,
          type: '精彩对局',
          replayId: 'SF6_20241130_003',
          description: '第二局的完美防守反击',
          timestamp: '2024-11-30T16:45:00.000Z',
          status: 'pending'
        }
      ]
    }
  },
  computed: {
    sortedSubmissions() {
      return [...this.submissions].sort((a, b) => 
        new Date(b.timestamp) - new Date(a.timestamp)
      )
    },
    pendingCount() {
      return this.submissions.filter(s => s.status === 'pending').length
    },
    todayCount() {
      const today = new Date().toDateString()
      return this.submissions.filter(s => 
        new Date(s.timestamp).toDateString() === today
      ).length
    }
  },
  methods: {
    formatTime(timestamp) {
      const date = new Date(timestamp)
      return date.toLocaleString('zh-CN', {
        year: 'numeric',
        month: '2-digit',
        day: '2-digit',
        hour: '2-digit',
        minute: '2-digit'
      })
    },
    getTypeClass(type) {
      const classes = {
        '超唐对局': 'type-super',
        '精彩对局': 'type-amazing',
        '上大师了!': 'type-master'
      }
      return classes[type] || 'type-default'
    },
    getStatusClass(status) {
      const classes = {
        'pending': 'status-pending',
        'approved': 'status-approved',
        'rejected': 'status-rejected'
      }
      return classes[status] || 'status-default'
    },
    getStatusText(status) {
      const texts = {
        'pending': '待处理',
        'approved': '已通过',
        'rejected': '已拒绝'
      }
      return texts[status] || '未知'
    }
  }
}
</script>

<style scoped>
.admin-header {
  text-align: center;
  margin-bottom: 40px;
  padding: 30px 0;
  border-bottom: 2px solid var(--accent-color);
}

.admin-title {
  font-size: 32px;
  color: var(--primary-color);
  margin-bottom: 10px;
  font-weight: 600;
}

.admin-subtitle {
  color: var(--light-text);
  font-size: 16px;
  margin: 0;
}

.stats-section {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
  gap: 20px;
  margin-bottom: 40px;
}

.stat-card {
  background-color: var(--background-color);
  padding: 20px;
  border-radius: var(--border-radius);
  text-align: center;
  border: 1px solid var(--border-color);
  box-shadow: var(--shadow);
}

.stat-number {
  font-size: 32px;
  font-weight: 600;
  color: var(--accent-color);
  margin-bottom: 8px;
}

.stat-label {
  color: var(--light-text);
  font-size: 14px;
}

.submissions-section {
  background-color: var(--background-color);
  padding: 30px;
  border-radius: var(--border-radius);
  border: 1px solid var(--border-color);
}

.empty-state {
  text-align: center;
  padding: 60px 20px;
  color: var(--light-text);
}

.empty-icon {
  font-size: 48px;
  margin-bottom: 20px;
}

.empty-hint {
  font-size: 14px;
  margin-top: 10px;
}

.submissions-list {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.submission-item {
  border: 1px solid var(--border-color);
  border-radius: var(--border-radius);
  padding: 20px;
  background-color: var(--light-background);
  transition: box-shadow 0.3s ease;
}

.submission-item:hover {
  box-shadow: var(--shadow);
}

.submission-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 15px;
}

.submission-type {
  padding: 6px 12px;
  border-radius: 20px;
  font-size: 12px;
  font-weight: 600;
  color: white;
}

.type-super {
  background-color: #e74c3c;
}

.type-amazing {
  background-color: #f39c12;
}

.type-master {
  background-color: #27ae60;
}

.type-default {
  background-color: var(--light-text);
}

.submission-time {
  color: var(--light-text);
  font-size: 12px;
}

.submission-content {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.submission-field {
  display: flex;
  align-items: center;
  gap: 10px;
}

.submission-field label {
  font-weight: 600;
  color: var(--primary-color);
  min-width: 80px;
  font-size: 14px;
}

.field-value {
  color: var(--text-color);
  font-size: 14px;
  flex: 1;
}

.status-badge {
  padding: 4px 8px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 600;
  color: white;
}

.status-pending {
  background-color: #f39c12;
}

.status-approved {
  background-color: #27ae60;
}

.status-rejected {
  background-color: #e74c3c;
}

.status-default {
  background-color: var(--light-text);
}

@media (max-width: 768px) {
  .admin-title {
    font-size: 24px;
  }
  
  .stats-section {
    grid-template-columns: 1fr;
  }
  
  .submission-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 10px;
  }
  
  .submission-field {
    flex-direction: column;
    align-items: flex-start;
    gap: 5px;
  }
  
  .submission-field label {
    min-width: auto;
  }
}
</style> 