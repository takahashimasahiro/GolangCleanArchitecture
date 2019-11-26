FROM mysql:5.7

# (1) パッケージアップデートとwgetインストール
RUN apt-get update
RUN apt-get install -y wget

# (2) entrykitのインストール
RUN wget https://github.com/progrium/entrykit/releases/download/v0.4.0/entrykit_0.4.0_linux_x86_64.tgz
RUN tar -xvzf entrykit_0.4.0_linux_x86_64.tgz
RUN rm entrykit_0.4.0_linux_x86_64.tgz
RUN mv entrykit /usr/local/bin/
RUN entrykit --symlink

# (3) スクリプトと各種設定ファイルのコピー
COPY add-server-id.sh /usr/local/bin/
COPY /mysql/mysql.conf.d/mysqld.cnf /mysql/mysql.conf.d/
COPY /mysql/conf.d/mysql.cnf /mysql/conf.d/
COPY prepare.sh /docker-entrypoint-initdb.d
COPY init-data.sh /usr/local/bin/
COPY db/setup.sql /db/setup.sql

# (4) スクリプトとmysqldの実行
ENTRYPOINT [ \
  "prehook", \
    "add-server-id.sh", \
    "--", \
  "docker-entrypoint.sh" \
]

CMD ["mysqld"]