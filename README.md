# Message-Board

### 在吴昊达学长留言板的基础下进行改动,并加入其他新功能

###### 功能实现

- [x] 注册，登陆，修改密码

- [x] 发表留言，修改留言，删除留言

- [x] 发表评论，修改评论，删除评论

- [x] 对入参进行检验，如用户名长度，密码长度

- [x] 文件上传

- [ ] 对留言与评论点赞

- [x] 密保功能

- [ ] 实现评论区的套娃

##### 建表代码

  ```mysql
  CREATE TABLE `user`
  (
      `Id`       BIGINT(20) NOT NULL AUTO_INCREMENT,
      `Name`     VARCHAR(20)  DEFAULT '',
      `Password` VARCHAR(20)  DEFAULT '123456',
      `Question` VARCHAR(255) DEFAULT NULL,
      `Answer`   VARCHAR(255) DEFAULT NULL,
      PRIMARY KEY (`Id`)
  ) ENGINE = InnoDB
    AUTO_INCREMENT = 1
    DEFAULT CHARSET = utf8mb4;
  ```

  ```mysql
  CREATE TABLE `post`
  (
      `Id`         BIGINT(20)             NOT NULL AUTO_INCREMENT,
      `Name`       VARCHAR(20)  DEFAULT '',
      `Context`    VARCHAR(255) DEFAULT NULL,
      `PostTime`   datetime     DEFAULT NULL,
      `UpdateTime` datetime     DEFAULT NULL,
      `Likes`      bigint       default 0 null,
      PRIMARY KEY (`Id`)
  ) ENGINE = InnoDB
    AUTO_INCREMENT = 1
    DEFAULT CHARSET = utf8mb4;
  ```

  ```mysql
  CREATE TABLE `Comment`
  (
      `Id`         BIGINT(20)             NOT NULL AUTO_INCREMENT,
      `PostId`     BIGINT(20)             NOT NULL,
      `Name`       VARCHAR(20)  DEFAULT '',
      `Context`    VARCHAR(255) DEFAULT NULL,
      `CommenTime` datetime     DEFAULT NULL,
      `Likes`      bigint       default 0 null,
      PRIMARY KEY (`Id`)
  ) ENGINE = InnoDB
    AUTO_INCREMENT = 1
    DEFAULT CHARSET = utf8mb4;
  ```

  