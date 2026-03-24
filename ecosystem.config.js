module.exports = {
  apps: [{
    name: 'githook',
    script: 'go',
    args: 'run main.go',
    cwd: '/root/githook',
    instances: 1,
    autorestart: true,
    watch: false,
    max_memory_restart: '1G',
    env: {
      REDIS_USR: '',
      REDIS_PW: '',
      REDIS_ADDR: '127.0.0.1:6379',
      SECRET: 'replace-me-with-a-random-string'
    },
    env_production: {
      NODE_ENV: 'production'
    },
    error_file: './logs/err.log',
    out_file: './logs/out.log',
    log_file: './logs/combined.log',
    time: true
  }]
};
