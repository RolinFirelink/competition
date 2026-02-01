<template>
  <div class="competition-list-container">
    <div class="header">
      <h1>赛事列表</h1>
    </div>
    <div class="competitions-wrapper">
      <div class="competition-item" v-for="competition in competitions" :key="competition.path">
        <router-link :to="competition.path" class="competition-link">
          <div class="competition-date">{{ competition.date }}</div>
          <div class="competition-title">{{ competition.title }}</div>
          <div class="competition-participants">参赛人数：{{ competition.participants }}人</div>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'CompetitionList',
  data() {
    return {
      competitions: [
        {
          path: '/20260131',
          date: '2026年1月31日',
          title: '1月31日杯赛',
          participants: 68
        }
        // 未来可以在这里添加更多赛事
      ]
    };
  }
};
</script>

<style scoped>
.competition-list-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 40px 20px;
  min-height: 100vh;
  background-color: #fafafa;
}

.header {
  text-align: center;
  margin-bottom: 40px;
}

.header h1 {
  font-size: 32px;
  font-weight: 600;
  color: #2c3e50;
  margin: 0;
}

.competitions-wrapper {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 24px;
}

.competition-item {
  background-color: #ffffff;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  transition: transform 0.3s ease, box-shadow 0.3s ease;
  overflow: hidden;
}

.competition-item:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.15);
}

.competition-link {
  display: block;
  padding: 24px;
  text-decoration: none;
  color: inherit;
}

.competition-date {
  font-size: 14px;
  color: #667eea;
  font-weight: 600;
  margin-bottom: 12px;
  text-transform: uppercase;
  letter-spacing: 1px;
}

.competition-title {
  font-size: 20px;
  font-weight: 600;
  color: #2c3e50;
  margin-bottom: 12px;
}

.competition-participants {
  font-size: 14px;
  color: #666666;
}

@media (max-width: 768px) {
  .competition-list-container {
    padding: 20px 10px;
  }

  .header h1 {
    font-size: 24px;
  }

  .competitions-wrapper {
    grid-template-columns: 1fr;
    gap: 16px;
  }

  .competition-link {
    padding: 20px;
  }

  .competition-title {
    font-size: 18px;
  }
}
</style>

