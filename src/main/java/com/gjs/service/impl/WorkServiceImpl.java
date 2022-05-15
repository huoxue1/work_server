package com.gjs.service.impl;

import com.gjs.entity.Work;
import com.gjs.dao.WorkDao;
import com.gjs.service.WorkService;
import org.springframework.stereotype.Service;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageImpl;
import org.springframework.data.domain.PageRequest;

import javax.annotation.Resource;
import java.util.List;

/**
 * (Work)表服务实现类
 *
 * @author makejava
 * @since 2022-05-15 13:45:27
 */
@Service("workService")
public class WorkServiceImpl implements WorkService {
    @Resource
    private WorkDao workDao;

    @Override
    public List<Work> queryAll() {
        return workDao.queryAll();
    }

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    @Override
    public Work queryById(Integer id) {
        return this.workDao.queryById(id);
    }

    /**
     * 分页查询
     *
     * @param work        筛选条件
     * @param pageRequest 分页对象
     * @return 查询结果
     */
    @Override
    public Page<Work> queryByPage(Work work, PageRequest pageRequest) {
        long total = this.workDao.count(work);
        return new PageImpl<>(this.workDao.queryAllByLimit(work, pageRequest), pageRequest, total);
    }

    /**
     * 新增数据
     *
     * @param work 实例对象
     * @return 实例对象
     */
    @Override
    public Work insert(Work work) {
        this.workDao.insert(work);
        return work;
    }

    /**
     * 修改数据
     *
     * @param work 实例对象
     * @return 实例对象
     */
    @Override
    public Work update(Work work) {
        this.workDao.update(work);
        return this.queryById(work.getId());
    }

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 是否成功
     */
    @Override
    public boolean deleteById(Integer id) {
        return this.workDao.deleteById(id) > 0;
    }
}
