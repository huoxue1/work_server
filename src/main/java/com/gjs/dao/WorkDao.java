package com.gjs.dao;

import com.gjs.entity.Work;
import org.apache.ibatis.annotations.Param;
import org.springframework.data.domain.Pageable;

import java.util.List;

/**
 * (Work)表数据库访问层
 *
 * @author makejava
 * @since 2022-05-15 13:45:24
 */
public interface WorkDao {



    List<Work> queryAll();

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    Work queryById(Integer id);

    /**
     * 查询指定行数据
     *
     * @param work     查询条件
     * @param pageable 分页对象
     * @return 对象列表
     */
    List<Work> queryAllByLimit(Work work, @Param("pageable") Pageable pageable);

    /**
     * 统计总行数
     *
     * @param work 查询条件
     * @return 总行数
     */
    long count(Work work);

    /**
     * 新增数据
     *
     * @param work 实例对象
     * @return 影响行数
     */
    int insert(Work work);

    /**
     * 批量新增数据（MyBatis原生foreach方法）
     *
     * @param entities List<Work> 实例对象列表
     * @return 影响行数
     */
    int insertBatch(@Param("entities") List<Work> entities);

    /**
     * 批量新增或按主键更新数据（MyBatis原生foreach方法）
     *
     * @param entities List<Work> 实例对象列表
     * @return 影响行数
     * @throws org.springframework.jdbc.BadSqlGrammarException 入参是空List的时候会抛SQL语句错误的异常，请自行校验入参
     */
    int insertOrUpdateBatch(@Param("entities") List<Work> entities);

    /**
     * 修改数据
     *
     * @param work 实例对象
     * @return 影响行数
     */
    int update(Work work);

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 影响行数
     */
    int deleteById(Integer id);

}

