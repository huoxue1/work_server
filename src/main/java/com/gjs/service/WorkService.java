package com.gjs.service;

import com.gjs.entity.Work;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;

import java.util.List;

/**
 * (Work)表服务接口
 *
 * @author makejava
 * @since 2022-05-15 13:45:26
 */
public interface WorkService {

    List<Work> queryAll();

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    Work queryById(Integer id);

    /**
     * 分页查询
     *
     * @param work        筛选条件
     * @param pageRequest 分页对象
     * @return 查询结果
     */
    Page<Work> queryByPage(Work work, PageRequest pageRequest);

    /**
     * 新增数据
     *
     * @param work 实例对象
     * @return 实例对象
     */
    Work insert(Work work);

    /**
     * 修改数据
     *
     * @param work 实例对象
     * @return 实例对象
     */
    Work update(Work work);

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 是否成功
     */
    boolean deleteById(Integer id);

}
