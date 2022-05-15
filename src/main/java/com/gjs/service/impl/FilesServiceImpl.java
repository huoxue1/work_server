package com.gjs.service.impl;

import com.gjs.config.Config;
import com.gjs.controller.FilesController;
import com.gjs.dao.WorkDao;
import com.gjs.entity.Files;
import com.gjs.dao.FilesDao;
import com.gjs.entity.Work;
import com.gjs.service.FilesService;
import com.gjs.utils.Resp;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageImpl;
import org.springframework.data.domain.PageRequest;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.web.multipart.MultipartFile;

import javax.annotation.Resource;
import java.io.File;
import java.io.IOException;
import java.util.Date;
import java.util.List;

/**
 * (Files)表服务实现类
 *
 * @author makejava
 * @since 2022-05-15 13:45:29
 */
@Transactional
@Service("filesService")
public class FilesServiceImpl implements FilesService {

    private static final org.slf4j.Logger logger = LoggerFactory.getLogger(FilesService.class);

    @Resource
    private FilesDao filesDao;

    @Resource
    private WorkDao workDao;

    @Resource
    Config config;

    @Override
    public boolean upload(Integer work_id, String token, MultipartFile multipartFile) throws IOException {
        String path = config.getStorage_path();
        Work work = workDao.queryById(work_id);
        Files files = new Files();
        files.setWorkId(work_id);
        files.setToken(token);
        files.setFileName(multipartFile.getOriginalFilename());
        long count = this.count(files);
        files.setSize((int) multipartFile.getSize());
        files.setUploadTime((int) new Date().getTime());
        files.setData("");

        if (count>0){
            this.update(files);
        }else {
            this.insert(files);
        }
        File file = new File(path + "/" + work.getName() + "/" + files.getFileName());
        multipartFile.transferTo(file);
        return true;

    }

    @Override
    public long count(Files files) {
        return filesDao.count(files);
    }

    @Override
    public boolean removeFile(Files file) {
        int i = filesDao.deleteById(file.getId());
        Work work = workDao.queryById(file.getWorkId());
        File f = new File(config.getStorage_path() + "/" + work.getName() + "/" + file.getFileName());
        f.deleteOnExit();
        logger.info("it will delete file in > "+f.getAbsolutePath());
        return i>0;
    }



    @Override
    public List<Files> queryByWorkId(Integer work_id) {
        return filesDao.queryByWorkId(work_id);
    }

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    @Override
    public Files queryById(Integer id) {
        return this.filesDao.queryById(id);
    }

    /**
     * 分页查询
     *
     * @param files       筛选条件
     * @param pageRequest 分页对象
     * @return 查询结果
     */
    @Override
    public Page<Files> queryByPage(Files files, PageRequest pageRequest) {
        long total = this.filesDao.count(files);
        return new PageImpl<>(this.filesDao.queryAllByLimit(files, pageRequest), pageRequest, total);
    }

    /**
     * 新增数据
     *
     * @param files 实例对象
     * @return 实例对象
     */
    @Override
    public Files insert(Files files) {
        this.filesDao.insert(files);
        return files;
    }

    /**
     * 修改数据
     *
     * @param files 实例对象
     * @return 实例对象
     */
    @Override
    public Files update(Files files) {
        this.filesDao.update(files);
        return this.queryById(files.getId());
    }

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 是否成功
     */
    @Override
    public boolean deleteById(Integer id) {
        return this.filesDao.deleteById(id) > 0;
    }
}
