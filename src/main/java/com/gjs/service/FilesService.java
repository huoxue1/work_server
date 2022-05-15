package com.gjs.service;

import com.gjs.entity.Files;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;
import java.util.List;

/**
 * (Files)表服务接口
 *
 * @author makejava
 * @since 2022-05-15 13:45:29
 */
public interface FilesService {

    boolean upload(Integer work_id, String token, MultipartFile multipartFile) throws IOException;

    long count(Files files);

    boolean removeFile(Files file);

    List<Files> queryByWorkId(Integer work_id);

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    Files queryById(Integer id);

    /**
     * 分页查询
     *
     * @param files       筛选条件
     * @param pageRequest 分页对象
     * @return 查询结果
     */
    Page<Files> queryByPage(Files files, PageRequest pageRequest);

    /**
     * 新增数据
     *
     * @param files 实例对象
     * @return 实例对象
     */
    Files insert(Files files);

    /**
     * 修改数据
     *
     * @param files 实例对象
     * @return 实例对象
     */
    Files update(Files files);

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 是否成功
     */
    boolean deleteById(Integer id);

}
