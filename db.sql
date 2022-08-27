CREATE TABLE ORDER
(
    uid serial,
    data json,
)

CREATE TABLE `order`
(
  `order_uid` varchar(128) NOT NULL ,
  `track_number` varchar(128) NOT NULL ,
  `entry` varchar(128) NOT NULL ,
  `locale` varchar(128) NOT NULL ,
  `internal_signature` varchar(128) NOT NULL ,
  `customer_id` varchar(128) NOT NULL ,
  `delivery_service` varchar(128) NOT NULL ,
  `shardkey` varchar(128) NOT NULL ,
  `sm_id` int NOT NULL ,
  `date_created` datetime NOT NULL ,
  `oof_shard` varchar(128) NOT NULL ,
  
) engine=innodb DEFAULT charset=utf8mb4;