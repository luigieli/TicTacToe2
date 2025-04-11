import styles from "./FrameBorder.module.css";

const LayoutBorder: React.FC = () => {
  return (
    <>
      <div className={`${styles.layoutBorder} ${styles.borderBottom}`} />
      <div className={`${styles.layoutBorder} ${styles.borderTop}`} />
    </>
  );
};

export default LayoutBorder;