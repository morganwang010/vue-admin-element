
		CREATE TABLE process (
  id SERIAL not Null,  -- 自增主键id
  creatorid INTEGER,     -- 创建者creator_id
  contractid INTEGER,    -- 合同contract_id
  created DATE,           -- 创建日期created
  amount NUMERIC,         -- 金额amount
  remarks TEXT, --备注
  PRIMARY key (id)
);
drop table process 

CREATE TABLE public.urls (
  id serial PRIMARY KEY,
  url character varying(255) DEFAULT NULL,
  description text DEFAULT NULL,
  keywords text DEFAULT NULL
);


select  * from product p ;
select  * from customer c where id = 243;
select * from user;

select *  from public.user where username = 'admin'

SELECT * FROM "public"."user" WHERE username = 'admin' ORDER BY "user"."id" LIMIT 1

CREATE TABLE public.user (
  id serial PRIMARY KEY,
  email character varying(30) DEFAULT NULL,
  password character varying(200) DEFAULT NULL,
  name character varying(10) DEFAULT NULL,
  status smallint DEFAULT '1',
  created bigint DEFAULT NULL,
  updated bigint DEFAULT NULL
);
COMMENT ON TABLE public.user IS '用户表';
COMMENT ON COLUMN public.user.id IS '编号';
COMMENT ON COLUMN public.user.email IS '邮箱';
COMMENT ON COLUMN public.user.password IS '密码';
COMMENT ON COLUMN public.user.name IS '名称';
COMMENT ON COLUMN public.user.status IS '状态，1-正常，2-注销';
COMMENT ON COLUMN public.user.created IS '创建时间';
COMMENT ON COLUMN public.user.updated IS '更新时间'
;

INSERT INTO public.user ("id", "email", "password", "name", "status", "created", "updated")
VALUES (29, '1655064994@qq.com', '$2a$10$62yO.fxSfNlstacxZfTtdO2uuR9YKG6hykuVTBIMc06CEJ3BWW/Ny', '', 1, 1671191625, 0);

CREATE TABLE public.contract (
  id bigserial NOT NULL,
  name character varying(200) DEFAULT NULL,
  amount numeric(10,2) DEFAULT NULL,
  begin_time character varying(50) DEFAULT NULL,
  over_time character varying(50) DEFAULT NULL,
  remarks character varying(80) DEFAULT NULL,
  cid bigint DEFAULT NULL,
  productlist jsonb DEFAULT NULL,
  status smallint DEFAULT NULL,
  creator bigint DEFAULT NULL,
  created bigint DEFAULT NULL,
  updated bigint DEFAULT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_creator FOREIGN KEY (creator) REFERENCES public.user (id)
);

COMMENT ON TABLE public.contract IS '合同表';
COMMENT ON COLUMN public.contract.id IS '编号';
COMMENT ON COLUMN public.contract.name IS '合同名称';
COMMENT ON COLUMN public.contract.amount IS '金额';
COMMENT ON COLUMN public.contract.begin_time IS '开始时间';
COMMENT ON COLUMN public.contract.over_time IS '结束时间';
COMMENT ON COLUMN public.contract.remarks IS '备注';
COMMENT ON COLUMN public.contract.cid IS '客户编号';
COMMENT ON COLUMN public.contract.productlist IS '产品编号和数量';
COMMENT ON COLUMN public.contract.status IS '状态';
COMMENT ON COLUMN public.contract.creator IS '创建人';
COMMENT ON COLUMN public.contract.created IS '创建时间';
COMMENT ON COLUMN public.contract.updated IS '更新时间';

INSERT INTO public.contract (id, name, amount, begin_time, over_time, remarks, cid, productlist, status, creator, created, updated) VALUES 
(20001, '电动车交易1', 89880.00, '2023-01-28', '2023-01-30', '无备注', 1, '[{"id": 1, "name": "电动车1", "type": 1, "unit": "台", "count": 10, "price": 1498, "total": 14980}, {"id": 2, "name": "电动车2", "type": 1, "unit": "台", "count": 20, "price": 1498, "total": 29960}, {"id": 3, "name": "电动车3", "type": 1, "unit": "台", "count": 30, "price": 1498, "total": 44940}]', 1, 29, 1674900672, 0);

INSERT INTO public.contract (id, name, amount, begin_time, over_time, remarks, cid, productlist, status, creator, created, updated) VALUES 



CREATE TABLE public.customer (
  id serial NOT NULL, -- 编号
  name varchar(50) DEFAULT NULL, -- 名称
  source char(15) DEFAULT NULL, -- 来源
  phone char(12) DEFAULT NULL, -- 手机号
  email char(20) DEFAULT NULL, -- 邮箱
  industry char(30) DEFAULT NULL, -- 所在行业
  level char(10) DEFAULT NULL, -- 级别
  remarks varchar(200) NOT NULL, -- 备注
  region char(80) DEFAULT NULL, -- 地区
  address varchar(255) DEFAULT NULL, -- 详细地址
  status smallint DEFAULT NULL, -- 成交状态
  creator bigint DEFAULT NULL, -- 创建人
  created bigint DEFAULT NULL, -- 创建时间
  updated bigint DEFAULT NULL, -- 更新时间
  PRIMARY KEY (id)
);



INSERT INTO public.customer (id, name, source, phone, email, industry, level, remarks, region, address, status, creator, created, updated)
VALUES (1, '北京文化有限公司1', '电话咨询', '13050803360', 'bjwenhua@gmail.com', '金融业', '重点客户', '', '北京市,朝阳区', '望京东园615号', 1, 29, 1674899545, 0);

INSERT INTO public.customer (id, name, source, phone, email, industry, level, remarks, region, address, status, creator, created, updated)
VALUES (2, '北京文化有限公司2', '线上注册', '13050803360', 'bjwenhua@gmail.com', '金融业', '重点客户', '', '北京市,朝阳区', '望京东园615号', 1, 29, 1674899545, 1674899808);

CREATE TABLE public.mail_config (
  id bigserial NOT NULL, -- 编号
  stmp char(50) DEFAULT NULL, -- ip地址或域名
  port int DEFAULT NULL, -- 端口号
  auth_code char(80) DEFAULT NULL, -- 授权码
  email char(80) NOT NULL, -- 邮箱账号
  status smallint DEFAULT NULL, -- 服务状态
  creator bigint DEFAULT NULL, -- 创建人
  created bigint DEFAULT NULL, -- 创建时间
  updated bigint DEFAULT NULL, -- 更新时间
  PRIMARY KEY (id)
);


INSERT INTO public.mail_config (id, stmp, port, auth_code, email, status, creator, created, updated)
VALUES (11, 'smtp.qq.com', 465, 'zrzxsebacrpfdaeg', '200300666@qq.com', 2, 29, 1674901189, 1674901237);



select * from urls u where keywords like '%公有云%'
update urls set status = 1;

CREATE TABLE public.product (
  id bigserial NOT NULL, -- 编号
  name varchar(50) DEFAULT NULL, -- 名称
  type smallint DEFAULT NULL, -- 类型
  unit char(5) DEFAULT NULL, -- 单位
  code varchar(80) DEFAULT NULL, -- 编码
  price numeric(10,2) DEFAULT NULL, -- 价格
  description varchar(200) DEFAULT NULL, -- 描述
  status smallint DEFAULT NULL, -- 状态，1-上架，2-下架
  creator bigint DEFAULT NULL, -- 创建人
  created bigint DEFAULT NULL, -- 创建时间
  updated bigint DEFAULT NULL, -- 更新时间
  PRIMARY KEY (id)
 
);

INSERT INTO public.product (id, name, type, unit, code, price, description, status, creator, created, updated)
VALUES (1, '电动车1', 1, '台', '004', 1498.00, '代驾折叠电动车电动自行车成人代步外卖锂电池小型轻便电瓶车迷你便携电单车 G2/汽车电芯-能量回收-6Ah约60km', 1, 29, 1671191995, 0);

INSERT INTO public.product (id, name, type, unit, code, price, description, status, creator, created, updated)
VALUES (2, '电动车2', 1, '台', '004', 1498.00, '代驾折叠电动车电动自行车成人代步外卖锂电池小型轻便电瓶车迷你便携电单车 G2/汽车电芯-能量回收-6Ah约60km', 1, 29, 1671191995, 0);




CREATE TABLE public.contract (
  id bigserial NOT NULL,
  name character varying(200) DEFAULT NULL,
  amount numeric(10,2) DEFAULT NULL,
  begin_time time DEFAULT NULL,
  over_time time DEFAULT NULL,
  remarks character varying(80) DEFAULT NULL,
  cid bigint DEFAULT NULL,
  productlist jsonb DEFAULT NULL,
  status smallint DEFAULT NULL,
  creator bigint DEFAULT NULL,
  created time DEFAULT NULL,
  updated time DEFAULT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_creator FOREIGN KEY (creator) REFERENCES public.user (id)
);


CREATE TABLE notice (
  id BIGSERIAL PRIMARY KEY, -- 编号
  content VARCHAR(200), -- 通知内容
  status SMALLINT, -- 状态，1-已读，2-未读
  creator BIGINT, -- 创建者
  created BIGINT, -- 创建时间
  updated BIGINT -- 更新时间
);

CREATE INDEX idx_creator ON notice (creator);