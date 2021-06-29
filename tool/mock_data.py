#!/usr/bin/env python
# -*- encoding: utf-8 -*-

import pymysql

def run():
    db = pymysql.connect(
            host='49.232.87.103',
            port=3306,
            user='zw',
            password='Zongwen123!@#',
            db='game_site',
            charset='utf8'
            )

    cursor = db.cursor()

    sql = "show create table auth;"
    cursor.execute(sql)
    print(cursor.fetchall())
    
    # 关闭数据库连接
    db.close()



run()

