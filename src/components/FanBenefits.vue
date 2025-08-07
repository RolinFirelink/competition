<template>
  <div class="section">
    <h2 class="section-title">粉丝福利</h2>
    
    <div class="benefits-content">
      <div class="benefits-intro">
        <div class="benefit-card">
          <div class="benefit-icon">🏆</div>
          <div class="benefit-text">
            <h3>粉丝福利</h3>
            <p>对所有第一次晋升大师的粉丝，将对局录像投稿给我，我将为他制作解说局！</p>
            <div class="benefit-requirements">
              <h4>投稿要求：</h4>
              <ul>
                <li>直接私信主包发送对局id</li>
                <li>并且还要发送你的街霸6Id给主包</li>
                <li>注意不符合格式的投稿不会被接受</li>
              </ul>
            </div>
          </div>
        </div>
      </div>

      <div class="benefits-note">
        <p><strong>注意：</strong>只有在私信中注明你的街霸6用户ID的投稿才算是合法投稿，不符合格式的投稿不会被接受。</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'FanBenefits'
}
</script>

<style scoped>
.benefits-content {
  display: flex;
  flex-direction: column;
  gap: 30px;
}

.benefits-intro {
  background-color: var(--light-background);
  padding: 30px;
  border-radius: var(--border-radius);
  border: 2px solid var(--accent-color);
}

.benefit-card {
  display: flex;
  align-items: flex-start;
  gap: 20px;
}

.benefit-icon {
  font-size: 48px;
  flex-shrink: 0;
}

.benefit-text {
  flex: 1;
}

.benefit-text h3 {
  color: var(--primary-color);
  font-size: 24px;
  margin-bottom: 15px;
  font-weight: 600;
}

.benefit-text p {
  color: var(--text-color);
  font-size: 16px;
  line-height: 1.6;
  margin-bottom: 20px;
}

.benefit-requirements {
  background-color: var(--background-color);
  padding: 20px;
  border-radius: var(--border-radius);
  border-left: 4px solid var(--accent-color);
}

.benefit-requirements h4 {
  color: var(--primary-color);
  font-size: 18px;
  margin-bottom: 15px;
  font-weight: 600;
}

.benefit-requirements ul {
  list-style: none;
  padding-left: 0;
}

.benefit-requirements li {
  margin-bottom: 8px;
  padding-left: 20px;
  position: relative;
  font-size: 14px;
  color: var(--text-color);
}

.benefit-requirements li::before {
  content: '✓';
  position: absolute;
  left: 0;
  top: 0;
  color: var(--accent-color);
  font-weight: bold;
}

.benefits-note {
  background-color: #fff3cd;
  border: 1px solid #ffeaa7;
  border-radius: var(--border-radius);
  padding: 20px;
  color: #856404;
}

.benefits-note p {
  margin: 0;
  font-size: 14px;
  line-height: 1.6;
}

@media (max-width: 768px) {
  .benefit-card {
    flex-direction: column;
    text-align: center;
  }
  
  .benefit-icon {
    align-self: center;
  }
  
  .benefits-intro {
    padding: 20px;
  }
}
</style> 