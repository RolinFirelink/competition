<template>
  <div class="section">
    <h2 class="section-title">特别鸣谢</h2>
    <div class="acknowledgments-grid">
      <div 
        v-for="person in acknowledgments" 
        :key="person.id" 
        class="acknowledgment-item"
      >
        <div class="person-avatar">
          <img 
            :src="person.avatar" 
            :alt="person.name"
            @error="handleImageError"
          />
        </div>
        <div class="person-info">
          <h4 class="person-name">{{ person.name }}</h4>
          <p class="person-message">{{ person.message }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Acknowledgments',
  data() {
    return {
      acknowledgments: [
        {
          id: 1,
          name: 'Gasly🔟🇫🇷',
          avatar: '/src/assets/images/community.jpg',
          message: '感谢Gasly🔟🇫🇷对我的赞助,包括资金上的帮助和支持我的PS5手柄，是你的赞助让我能够坚持创作优质内容,真的非常感谢！'
        }
        // 后续可以通过硬编码添加更多鸣谢对象
      ]
    }
  },
  methods: {
    handleImageError(event) {
      // 如果头像加载失败，使用默认头像
      event.target.src = 'data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iODAiIGhlaWdodD0iODAiIHZpZXdCb3g9IjAgMCA4MCA4MCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KPHJlY3Qgd2lkdGg9IjgwIiBoZWlnaHQ9IjgwIiByeD0iNDAiIGZpbGw9IiMzNDk4ZGIiLz4KPGNpcmNsZSBjeD0iNDAiIGN5PSIzMCIgcj0iMTUiIGZpbGw9IndoaXRlIi8+CjxwYXRoIGQ9Ik0xNSA2NUMyMCA1NSA0MCA1NSA0NSA2NUgxNVoiIGZpbGw9IndoaXRlIi8+Cjwvc3ZnPgo='
    }
  }
}
</script>

<style scoped>
.acknowledgments-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 20px;
}

.acknowledgment-item {
  display: flex;
  align-items: center;
  gap: 15px;
  padding: 20px;
  border: 1px solid var(--border-color);
  border-radius: var(--border-radius);
  background-color: var(--light-background);
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.acknowledgment-item:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow);
}

.person-avatar {
  flex-shrink: 0;
}

.person-avatar img {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid var(--accent-color);
}

.person-info {
  flex: 1;
}

.person-name {
  font-size: 18px;
  font-weight: 600;
  color: var(--primary-color);
  margin-bottom: 8px;
}

.person-message {
  font-size: 14px;
  color: var(--text-color);
  line-height: 1.6;
  margin: 0;
}

@media (max-width: 768px) {
  .acknowledgments-grid {
    grid-template-columns: 1fr;
  }
  
  .acknowledgment-item {
    flex-direction: column;
    text-align: center;
  }
}
</style> 