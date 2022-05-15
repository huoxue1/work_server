package com.gjs.dao;

import com.gjs.entity.Files;
import org.apache.ibatis.annotations.Param;
import org.springframework.data.domain.Pageable;

import java.util.List;

/**
 * (Files)表数据库访问层
 *
 * @author makejava
 * @since 2022-05-15 13:45:27
 */
public interface FilesDao {

    List<Files> queryByWorkId(@Param("work_id") Integer work_id);

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    Files queryById(Integer id);

    /**
     * 查询指定行数据
     *
     * @param files    查询条件
     * @param pageable 分页对象
     * @return 对象列表
     */
    List<Files> queryAllByLimit(Files files, @Param("pageable") Pageable pageable);

    /**
     * 统计总行数
     *
     * @param files 查询条件
     * @return 总行数
     */
    long count(Files files);

    /**
     * 新增数据
     *
     * @param files 实例对象
     * @return 影响行数
     */
    int insert(Files files);

    /**
     * 批量新增数据（MyBatis原生foreach方法）
     *
     * @param entities List<Files> 实例对象列表
     * @return 影响行数
     */
    int insertBatch(@Param("entities") List<Files> entities);

    /**
     * 批量新增或按主键更新数据（MyBatis原生foreach方法）
     *
     * @param entities List<Files> 实例对象列表
     * @return 影响行数
     * @throws org.springframework.jdbc.BadSqlGrammarException 入参是空List的时候会抛SQL语句错误的异常，请自行校验入参
     */
    int insertOrUpdateBatch(@Param("entities") List<Files> entities);

    /**
     * 修改数据
     *
     * @param files 实例对象
     * @return 影响行数
     */
    int update(Files files);

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 影响行数
     */
    int deleteById(Integer id);

}

