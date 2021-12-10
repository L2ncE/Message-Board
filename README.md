# Message-Board

#### 在吴昊达学长留言板的基础下进行改动,并加入其他新功能

***LV3思路***  通过建一张NestedReply表来储存嵌套回复，第一次对一个CommentId回复为第一个楼中楼回复，若再一次回复则会进行嵌套。

通过使用这段代码：

```mysql
`update nestedReply set context=CONCAT_WS('/ ', context, ?) where CommentId = ?`
```

来实现嵌套，效果则是每一次对同一个评论进行楼中楼评论时，会保留楼中楼的内容并用"/"相隔开。

##### 功能实现

- [x] 注册，登陆，修改密码
- [x] 发表留言，修改留言，删除留言
- [x] 发表评论，修改评论，删除评论
- [x] 对入参进行检验，如用户名长度，密码长度
- [x] 文件上传
- [x] 管理员功能（能够删除所有发言）
- [x] 对留言、评论和楼中楼点赞
- [x] 密保功能
- [x] 匿名评论功能
- [x] 实现评论区的套娃
- [x] 删除套娃

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

  ```mysql
CREATE TABLE `NestedReply`
(
    `Id`        BIGINT(20)             NOT NULL AUTO_INCREMENT,
    `CommentId` BIGINT(20)             NOT NULL,
    `Name`      VARCHAR(20)  DEFAULT '',
    `Context`   VARCHAR(255) DEFAULT NULL,
    `ReplyTime` datetime     DEFAULT NULL,
    `Likes`     bigint       default 0 null,
    PRIMARY KEY (`Id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4;
  ```