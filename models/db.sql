-- mysql workbench forward engineering

set @old_unique_checks=@@unique_checks, unique_checks=0;
set @old_foreign_key_checks=@@foreign_key_checks, foreign_key_checks=0;
set @old_sql_mode=@@sql_mode, sql_mode='traditional,allow_invalid_dates';

-- -----------------------------------------------------
-- schema mydb
-- -----------------------------------------------------
-- -----------------------------------------------------
-- schema evfs_browser
-- -----------------------------------------------------

-- -----------------------------------------------------
-- schema evfs_browser
-- -----------------------------------------------------
create schema if not exists `evfs_browser` default character set utf8mb4 ;
use `evfs_browser` ;

-- -----------------------------------------------------
-- table `main_account`
-- -----------------------------------------------------
create table if not exists `main_account` (
  `account_address` varchar(100) not null,
  `nonce` int(11) null default null,
  `balance` varchar(45) null default null,
  `create_time` timestamp null default null,
  `need_refresh` int(11) null default '0' comment '是否需要刷新 0 不需要。1需要',
  `max` varchar(100) null default null,
  `accept_max` varchar(100) null default null,
  `accept_limit` varchar(100) null default null,
  `is_union` int(11) null default '0' comment '是否是联合账户，默认不是',
  primary key (`account_address`))
engine = innodb
default character set = utf8mb4;


-- -----------------------------------------------------
-- table `main_account_crypto_balance`
-- -----------------------------------------------------
create table if not exists `main_account_crypto_balance` (
  `crypto_token_hash` varchar(100) not null comment '加密token的标识。对应的主链上的crypt token hash。',
  `crypto_name` varchar(50) not null comment '加密token的定制。外键。',
  `account_address` varchar(100) not null comment '该token的拥有者地址。',
  `transaction_hash` varchar(100) not null comment 'token的拥有者通过该笔交易获取的',
  `token_code` varchar(45) null default null comment 'token的基本信息。编号。',
  `token_name` varchar(45) null default null comment 'token的基本信息。名称。',
  `nonce` int(11) null default null comment 'token的基本信息。被交易的次数。',
  `owner_time` timestamp null default null comment '拥有者何时获得的该token。',
  `create_time` timestamp null default null comment 'token的创建时间。',
  `block_hash` varchar(100) null default null comment '交易所在的区块hash',
  `token_data` varchar(2000) null default null,
  `need_refresh` int(11) null default '0' comment '是否需要刷新 0 不需要。1需要',
  primary key (`crypto_token_hash`))
engine = innodb
default character set = utf8mb4;


-- -----------------------------------------------------
-- table `main_account_token_balance`
-- -----------------------------------------------------
create table if not exists `main_account_token_balance` (
  `token_balance_id` int(11) not null auto_increment,
  `account_address` varchar(100) not null,
  `token_name` varchar(50) not null,
  `balance` varchar(50) null default null comment '可用金额',
  `need_refresh` int(11) null default '0' comment '是否需要刷新 0 不需要。1需要',
  `locked_balance` varchar(100) null default null comment '锁定金额',
  `freeze_balance` varchar(100) null default null comment '冻结金额',
  primary key (`token_balance_id`))
engine = innodb
default character set = utf8mb4;


-- -----------------------------------------------------
-- table `main_biz_contract`
-- -----------------------------------------------------
create table if not exists `main_biz_contract` (
  `contract_id` int(11) not null auto_increment,
  `storage_biz_id` varchar(100) null default null,
  `contract_name` varchar(200) null default null,
  `pub_time` datetime null default null,
  `is_active` int(11) null default null,
  `active_time` datetime null default null,
  primary key (`contract_id`))
engine = innodb
auto_increment = 2
default character set = utf8mb4;


-- -----------------------------------------------------
-- table `main_biz_data_type`
-- -----------------------------------------------------
create table if not exists `main_biz_data_type` (
  `data_type_id` int(11) not null auto_increment,
  `main_biz_contract_id` int(11) null default null,
  `contract_name` varchar(200) null default null,
  `type_name` varchar(200) null default null,
  `data_form` varchar(200) null default null,
  primary key (`data_type_id`))
engine = innodb
auto_increment = 3
default character set = utf8mb4;


-- -----------------------------------------------------
-- table `main_block`
-- -----------------------------------------------------
create table if not exists `main_block` (
  `block_id` int(11) not null auto_increment,
  `block_hash` varchar(100) not null comment '区块hash',
  `parent_block_hash` varchar(100) not null comment '上一个区块hash',
  `minner_account_address` varchar(100) not null comment 'block的出块节点的账户地址。区块的奖励会发给该地址。',
  `height` bigint(20) not null comment '区块高度。',
  `block_size` int(11) null default null,
  `timestamp` timestamp null default null comment '区块时间戳。秒',
  `state_root` varchar(100) null default null comment '状态树的根',
  `receipt_root` varchar(100) null default null comment '交易执行结果的根',
  `reward` varchar(50) null default null comment '区块奖励金额',
  `extra_data` varchar(2000) null default null comment '区块的扩展信息',
  `transaction_count` int(11) null default null comment '区块中包含的交易的数量',
  `created` datetime null default current_timestamp,
  `updated` datetime null default current_timestamp,
  primary key (`block_id`),
  index `idx_main_block_height` (`height` asc))
engine = innodb
auto_increment = 920892
default character set = utf8mb4;


-- -----------------------------------------------------
-- table `main_chaingroup`
-- -----------------------------------------------------
create table if not exists `main_chaingroup` (
  `group_id` varchar(100) not null,
  `name` varchar(45) null default null,
  `rule` int(11) null default null,
  primary key (`group_id`))
engine = innodb
default character set = utf8mb4;


-- -----------------------------------------------------
-- table `main_accountnode` 记账节点
-- -----------------------------------------------------
create table if not exists `main_accountnode` (
  `chainnode_id` varchar(100) not null,
  `main_org_org_id` varchar(100) not null,
  `chainnode_name` varchar(45) null default null,
  `org_name` varchar(100) null default null,
  `capacity_file_size` bigint(20) null default null,
  `capacity_data_size` bigint(20) null default null,
  `cpu` varchar(100) null default null,
  `memory` varchar(100) null default null,
  `disk` varchar(100) null default null,
  `bandwidth` varchar(100) null,
  primary key (`chainnode_id`))
engine = innodb
default character set = utf8mb4
comment = '记账节点';


-- -----------------------------------------------------
-- table `main_client`
-- -----------------------------------------------------
create table if not exists `main_client` (
  `client_id` varchar(100) not null,
  `client_name` varchar(45) null default null,
  `main_org_org_id` varchar(100) not null,
  primary key (`client_id`))
engine = innodb
default character set = utf8mb4;


-- -----------------------------------------------------
-- table `main_contract`
-- -----------------------------------------------------
create table if not exists `main_contract` (
  `contract_id` int(11) not null auto_increment,
  `contract_address` varchar(100) not null,
  `account_address` varchar(100) not null,
  `transaction_hash` varchar(100) not null,
  `nonce` int(11) null default null,
  `code_bin` varchar(4000) null default null,
  `code` varchar(4000) null default null,
  `create_time` timestamp null default null,
  `block_hash` varchar(100) null default null,
  `need_refresh` int(11) null default '0',
  `block_height` bigint(20) null default null,
  `timestamp` timestamp null default null,
  primary key (`contract_id`))
engine = innodb
default character set = utf8mb4;


-- -----------------------------------------------------
-- table `main_crypto`
-- -----------------------------------------------------
create table if not exists `main_crypto` (
  `crypto_id` int(11) not null auto_increment,
  `crypto_name` varchar(50) not null,
  `account_address` varchar(100) not null,
  `transaction_hash` varchar(100) not null,
  `total_count` int(11) null default null,
  `current_count` int(11) null default null,
  `create_time` timestamp null default null,
  `holders` int(11) null default null,
  `transfers` int(11) null default null,
  `block_hash` varchar(100) null default null,
  `block_height` bigint(20) null default null,
  `need_refresh` int(11) null default '0' comment '是否需要更新。重新统计holders和transfers',
  primary key (`crypto_id`))
engine = innodb
default character set = utf8mb4;


-- -----------------------------------------------------
-- table `main_org`
-- -----------------------------------------------------
create table if not exists `main_org` (
  `org_id` varchar(100) not null,
  `org_name` varchar(45) null default null,
  primary key (`org_id`))
engine = innodb
default character set = utf8mb4;


-- -----------------------------------------------------
-- table `main_resourcenode`
-- -----------------------------------------------------
create table if not exists `main_resourcenode` (
  `resourcenode_id` varchar(100) not null,
  `main_org_org_id` varchar(100) not null,
  `main_storage_storage_id` varchar(100) not null,
  `capacity_file_size` bigint(20) null default null,
  `capacity_data_size` bigint(20) null default null,
  `used_file_size` bigint(20) null default null,
  `used_data_size` bigint(20) null default null,
  `file_count` int(11) null default null,
  `data_count` int(11) null default null,
  `cpu` varchar(100) null default null,
  `memory` varchar(100) null default null,
  `disk` varchar(100) null default null,
  `nic` varchar(100) null default null,
  `file_size` bigint(20) null default null,
  `data_size` bigint(20) null default null,
  primary key (`resourcenode_id`))
engine = innodb
default character set = utf8mb4
comment = '资源节点';


-- -----------------------------------------------------
-- table `main_storage`
-- -----------------------------------------------------
create table if not exists `main_storage` (
  `storage_id` varchar(100) not null,
  `storage_name` varchar(200) null default null,
  `capacity_file_size` bigint(20) null default null,
  `capacity_data_size` bigint(20) null default null,
  `used_file_size` bigint(20) null default null,
  `used_data_size` bigint(20) null default null,
  `file_count` int(11) null default null,
  `data_count` int(11) null default null,
  `org_num` int(11) null default null,
  `sys_num` int(11) null default null,
  `user_num` int(11) null default null,
  `out_chain_num` int(11) null default null,
  `node_num` int(11) null default null,
  `file_size` bigint(20) null default null,
  `data_size` bigint(20) null default null,
  `client_num` int(11) null default null,
  `committee_num` int(11) null default null,
  `rule` varchar(45) null,
  `main_org_org_id` varchar(45) null,
  primary key (`storage_id`))
engine = innodb
default character set = utf8mb4
comment = '数据存管域';


-- -----------------------------------------------------
-- table `main_biz`
-- -----------------------------------------------------
create table if not exists `main_biz` (
  `biz_id` varchar(100) not null,
  `main_storage_storage_id` varchar(100) not null,
  `biz_name` varchar(200) null default null,
  `main_storage_bizcol` varchar(45) null default null,
  `used_file_size` int(11) null default null,
  `used_data_size` int(11) null default null,
  `file_count` int(11) null default null,
  `data_count` int(11) null default null,
  `org_num` int(11) null default null,
  `sys_num` int(11) null default null,
  `user_num` int(11) null default null,
  `out_chain_num` int(11) null default null,
  `biz_data_type_count` int(11) null default null,
  `biz_contract_count` int(11) null default null,
  `file_size` int(11) null default null,
  `data_size` int(11) null default null,
  primary key (`biz_id`))
engine = innodb
default character set = utf8mb4
comment = '业务域';


-- -----------------------------------------------------
-- table `main_biz_member`
-- -----------------------------------------------------
create table if not exists `main_biz_member` (
  `committee_id` int unsigned not null auto_increment,
  `main_biz_biz_id` varchar(100) not null,
  `address` varchar(100) null,
  `name` varchar(45) null default null,
  `join_time` datetime null default null,
  primary key (`committee_id`))
engine = innodb
default character set = utf8mb4;


-- -----------------------------------------------------
-- table `main_storage_data`
-- -----------------------------------------------------
create table if not exists `main_storage_data` (
  `data_id` varchar(100) not null,
  `data_size` bigint(20) null default null,
  `main_system_system_id` varchar(100) not null,
  `main_storage_biz_biz_id` varchar(100) not null,
  `main_resourcenode_resourcenode_id` varchar(100) not null,
  `upload_time` datetime null default null comment '上链时间',
  primary key (`data_id`))
engine = innodb
default character set = utf8mb4;


-- -----------------------------------------------------
-- table `main_storage_file`
-- -----------------------------------------------------
create table if not exists `main_storage_file` (
  `file_id` varchar(100) not null,
  `file_name` varchar(1000) null default null,
  `file_size` bigint(20) null default null,
  `main_system_system_id` varchar(100) not null,
  `main_storage_biz_biz_id` varchar(100) not null,
  `main_resourcenode_resourcenode_id` varchar(100) not null,
  `file_type` int(11) null default null,
  `copy_count` int(11) null default null,
  `upload_time` datetime null default null comment '上链时间',
  primary key (`file_id`))
engine = innodb
default character set = utf8mb4;


-- -----------------------------------------------------
-- table `main_storage_member`
-- -----------------------------------------------------
create table if not exists `main_storage_member` (
  `member_id` int not null auto_increment,
  `main_storage_storage_id` varchar(100) not null,
  `address` varchar(100) null,
  `name` varchar(45) null default null,
  `join_time` datetime null default null,
  primary key (`member_id`))
engine = innodb
default character set = utf8mb4;


-- -----------------------------------------------------
-- table `main_system`
-- -----------------------------------------------------
create table if not exists `main_system` (
  `system_id` varchar(100) not null,
  `system_name` varchar(1000) null default null,
  `main_org_org_id` varchar(100) not null,
  `used_file_size` int(11) null default null,
  `used_data_size` int(11) null default null,
  `file_count` int(11) null default null,
  `data_count` int(11) null default null,
  `user_count` int(11) null default null,
  `file_size` int(11) null default null,
  `data_size` int(11) null default null,
  `main_storage_storage_id` varchar(100) null,
  `main_biz_biz_id` varchar(100) null,
  primary key (`system_id`))
engine = innodb
default character set = utf8mb4
comment = '业务系统';


-- -----------------------------------------------------
-- table `main_token`
-- -----------------------------------------------------
create table if not exists `main_token` (
  `token_id` int(11) not null auto_increment,
  `token_name` varchar(50) not null,
  `account_address` varchar(100) not null,
  `transaction_hash` varchar(100) not null,
  `total_amount` varchar(50) null default null,
  `create_time` timestamp null default null,
  `holders` int(11) null default null,
  `transfers` int(11) null default null,
  `block_hash` varchar(100) null default null,
  `block_height` bigint(20) null default null,
  `need_refresh` int(11) null default '0' comment '是否需要更新。重新统计holders和transfers',
  primary key (`token_id`))
engine = innodb
default character set = utf8mb4;


-- -----------------------------------------------------
-- table `main_transaction`
-- -----------------------------------------------------
create table if not exists `main_transaction` (
  `transaction_id` int(11) not null auto_increment,
  `transaction_hash` varchar(100) not null,
  `block_height` bigint(20) not null,
  `block_hash` varchar(100) not null,
  `peer_id` varchar(100) not null,
  `status` varchar(1) null default null,
  `result` varchar(2000) null default null,
  `timestamp` timestamp null default null,
  `inner_codetype` int(11) null default null,
  `code_data` varchar(4000) null default null,
  `extra_data` varchar(4000) null default null,
  `fee_hi` varchar(45) null default null,
  `fee_low` varchar(45) null default null,
  primary key (`transaction_id`),
  index `idx_main_transaction_main_block_block_height` (`block_height` asc))
engine = innodb
default character set = utf8mb4;


-- -----------------------------------------------------
-- table `main_transaction_call_contract`
-- -----------------------------------------------------
create table if not exists `main_transaction_call_contract` (
  `call_contract_id` int(11) not null auto_increment,
  `contract_address` varchar(100) not null,
  `transaction_hash` varchar(100) not null,
  `account_address` varchar(100) not null,
  `data` varchar(4000) null default null,
  `amount` varchar(50) null default null,
  `timestamp` timestamp null default null,
  `block_hash` varchar(100) null default null,
  primary key (`call_contract_id`))
engine = innodb
default character set = utf8mb4;


-- -----------------------------------------------------
-- table `main_transaction_crypto`
-- -----------------------------------------------------
create table if not exists `main_transaction_crypto` (
  `transaction_crypto_id` int(11) not null auto_increment,
  `transaction_hash` varchar(100) not null,
  `total` int(11) null default null,
  `current_total` int(11) null default null,
  `symbol` varchar(45) null default null,
  `name` varchar(4000) null default null,
  `code` varchar(4000) null default null,
  `props` varchar(4000) null default null,
  `timestamp` timestamp null default null,
  `block_hash` varchar(100) null default null,
  primary key (`transaction_crypto_id`))
engine = innodb
default character set = utf8mb4;


-- -----------------------------------------------------
-- table `main_transaction_token`
-- -----------------------------------------------------
create table if not exists `main_transaction_token` (
  `transaction_token_id` int(11) not null auto_increment,
  `block_hash` varchar(100) not null,
  `transaction_hash` varchar(100) not null,
  `account_address` varchar(100) not null,
  `token_name` varchar(45) null default null,
  `before_opt` varchar(45) null default null,
  `after_opt` varchar(45) null default null,
  `amount` varchar(45) null default null,
  `opt_type` int(11) null default null,
  `timestamp` timestamp null default null,
  `create_time` timestamp null default null,
  primary key (`transaction_token_id`))
engine = innodb
default character set = utf8mb4;


-- -----------------------------------------------------
-- table `main_transaction_transfer`
-- -----------------------------------------------------
create table if not exists `main_transaction_transfer` (
  `transaction_transfer_id` varchar(100) not null,
  `account_address` varchar(100) not null,
  `transaction_hash` varchar(100) not null,
  `token_name` varchar(50) null default null,
  `crypto_name` varchar(50) null default null,
  `block_hash` varchar(100) null default null,
  `transfer_type` varchar(10) null default null comment 'in / out',
  `amount` varchar(50) null default null,
  `token_amount` varchar(50) null default null,
  primary key (`transaction_transfer_id`))
engine = innodb
default character set = utf8mb4;


-- -----------------------------------------------------
-- table `main_transaction_transfer_crypto`
-- -----------------------------------------------------
create table if not exists `main_transaction_transfer_crypto` (
  `transaction_transfer_crypto_id` int(11) not null,
  `transaction_transfer_id` varchar(100) not null,
  `crypto_token_hash` varchar(100) not null,
  primary key (`transaction_transfer_crypto_id`))
engine = innodb
default character set = utf8mb4;


-- -----------------------------------------------------
-- table `main_union_account_rel`
-- -----------------------------------------------------
create table if not exists `main_union_account_rel` (
  `union_account_rel_id` int(11) not null,
  `sub_account_address` varchar(100) not null,
  `union_account_address` varchar(100) not null,
  primary key (`union_account_rel_id`))
engine = innodb
default character set = utf8mb4
comment = '联合账户的子账户关联表';


-- -----------------------------------------------------
-- table `main_user`
-- -----------------------------------------------------
create table if not exists `main_user` (
  `user_id` varchar(100) not null,
  `main_system_system_id` varchar(100) not null,
  `user_name` varchar(45) null default null,
  primary key (`user_id`))
engine = innodb
default character set = utf8mb4
comment = '数据上链用户';


-- -----------------------------------------------------
-- table `network_peer`
-- -----------------------------------------------------
create table if not exists `network_peer` (
  `peer_id` varchar(100) not null comment '节点id。来源于节点地址\\n因为bcuid在每次节点重启后会变化',
  `peer_address` varchar(100) null default null comment '地址。对应的节点的账户地址。',
  `peer_name` varchar(100) null default null comment '节点名称。',
  `status` varchar(45) null default null comment '节点状态。',
  `last_active_time` timestamp null default null comment '暂时无用',
  `last_alive_time` timestamp null default null comment '最后一次激活时间',
  `type` varchar(45) null default null comment '类型',
  `peer_idx` int(11) null default null comment 'node_idx',
  `peer_uri` varchar(200) null default null comment 'uri',
  primary key (`peer_id`))
engine = innodb
default character set = utf8mb4
comment = '节点信息表';


-- -----------------------------------------------------
-- table `rel_biz_consumer`
-- -----------------------------------------------------
create table if not exists `rel_biz_consumer` (
  `consumer_id` int(11) not null auto_increment,
  `main_storage_biz_biz_id` varchar(100) not null,
  `main_system_system_id` varchar(100) not null,
  primary key (`consumer_id`))
engine = innodb
default character set = utf8mb4;


-- -----------------------------------------------------
-- table `rel_biz_producer`
-- -----------------------------------------------------
create table if not exists `rel_biz_producer` (
  `producer_id` int(11) not null auto_increment,
  `main_storage_biz_biz_id` varchar(100) not null,
  `main_system_system_id` varchar(100) not null,
  primary key (`producer_id`))
engine = innodb
auto_increment = 3
default character set = utf8mb4;


-- -----------------------------------------------------
-- table `rel_org_storage`
-- -----------------------------------------------------
create table if not exists `rel_org_storage` (
  `rel_id` int(11) not null,
  `main_org_org_id` varchar(100) not null,
  `main_storage_storage_id` varchar(100) not null,
  primary key (`rel_id`))
engine = innodb
default character set = utf8mb4;


-- -----------------------------------------------------
-- table `sys_dict`
-- -----------------------------------------------------
create table if not exists `sys_dict` (
  `dict_key` varchar(50) not null,
  `dict_value` varchar(255) null default null,
  `dict_init` int(11) null default '0',
  primary key (`dict_key`))
engine = innodb
default character set = utf8mb4;


-- -----------------------------------------------------
-- table `main_chaingroup_member`
-- -----------------------------------------------------
create table if not exists `main_chaingroup_member` (
  `chaincommittee_id` int(11) not null,
  `main_chaingroup_group_id` varchar(100) not null,
  `member_address` varchar(100) null,
  `main_chaingroup_group_id` varchar(200) null default null,
  `join_time` datetime null default null,
  primary key (`chaincommittee_id`))
engine = innodb
default character set = utf8mb4;


-- -----------------------------------------------------
-- table `main_clientnode` 前置节点
-- -----------------------------------------------------
create table if not exists `main_clientnode` (
  `chainnode_id` varchar(100) not null,
  `main_org_org_id` varchar(100) not null,
  `chainnode_name` varchar(45) null default null,
  `org_name` varchar(100) null default null,
  `capacity_file_size` bigint(20) null default null,
  `capacity_data_size` bigint(20) null default null,
  `cpu` varchar(100) null default null,
  `memory` varchar(100) null default null,
  `disk` varchar(100) null default null,
  `nic` varchar(100) null default null,
  `bandwidth` varchar(100) null default null,
  primary key (`chainnode_id`))
engine = innodb
default character set = utf8mb4
comment = '前置节点';


-- -----------------------------------------------------
-- table `main_syncnode` 只读节点
-- -----------------------------------------------------
create table if not exists `main_syncnode` (
  `chainnode_id` varchar(100) not null,
  `main_org_org_id` varchar(100) not null,
  `chainnode_name` varchar(45) null default null,
  `org_name` varchar(100) null default null,
  `capacity_file_size` bigint(20) null default null,
  `capacity_data_size` bigint(20) null default null,
  `cpu` varchar(100) null default null,
  `memory` varchar(100) null default null,
  `disk` varchar(100) null default null,
  `bandwidth` varchar(100) null default null,
  primary key (`chainnode_id`))
engine = innodb
default character set = utf8mb4
comment = '前置节点';


set sql_mode=@old_sql_mode;
set foreign_key_checks=@old_foreign_key_checks;
set unique_checks=@old_unique_checks;


truncate table main_Block;
truncate main_transaction;
truncate `main_account_transaction`;
truncate main_account;
truncate evfs_biz;
truncate evfs_biz_member;
truncate evfs_chaingroup;
truncate evfs_node;
truncate evfs_org;
truncate evfs_resourcenode;
truncate evfs_storage;
truncate evfs_storage_member;
truncate evfs_system;